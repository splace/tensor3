package tensor3

type VectorRefs []*Vector

func NewVectorRefs(cs ...Scalar) (vs VectorRefs) {
	vs = make(VectorRefs, (len(cs)+2)/3)
	for i := range vs {
		v := NewVector(cs[i*3:]...)
		vs[i] = &v
	}
	return
}

// make a new VectorRefs that references Vector's at the provided indexes in the provided Vectors.
// Notice: indexes are uint and considered out-of-context, not an internal slice index, 1 is used for the first item, 0 used as a side-channel. (possibly for an error indicator).
func NewVectorRefsFromIndexes(cs Vectors, indexes ...uint) (vs VectorRefs) {
	if len(indexes) == 0 {
		vs = make(VectorRefs, len(cs))
		for i := range cs {
			vs[i] = &cs[i]
		}
	} else {
		vs = make(VectorRefs, len(indexes))
		for i := range vs {
			vs[i] = &cs[indexes[i]-1]
		}
	}
	return
}

// rebases, a number of VectorRefs, to point into a newly created Vectors.
func NewVectorsFromVectorRefs(vss ...VectorRefs) Vectors {
	m := make(map[*Vector]uint)
	for _, vs := range vss {
		for _, v := range vs {
			if _, ok := m[v]; !ok {
				m[v] = uint(len(m))
			}
		}
	}
	nv := make(Vectors, len(m))
	for vr, index := range m {
		nv[index] = *vr
	}
	for _, vs := range vss {
		for i, vr := range vs {
			vs[i] = &nv[m[vr]]
		}
	}
	return nv
}

// return a slice of indexes, the same length as the VectorRefs, to the items this references in the provided Vectors.
// (Notice: this package uses 1 for the first item.)
// an index of zero indicates reference not found.
func (vsr VectorRefs) Indexes(vs Vectors) (is []uint) {
	// TODO find index using unsafe pointer offset, massively faster.
	is = make([]uint, len(vsr))
	for ir, vr := range vsr {
		is[ir] = vs.Index(vr)
	}
	return
}

// make a slice of Vectors from the values referenced by this VectorRefs.
func (vsr VectorRefs) Dereference() (vs Vectors) {
	vs = make(Vectors, len(vsr))
	for i := range vs {
		vs[i] = *vsr[i]
	}
	return
}

// overwrite the references in a VectorRefs to refer to the Vector's.
// for upto the greater length of the two slices.
func (vs Vectors) Reference(vsr VectorRefs) {
	if len(vs) > len(vsr) {
		for i := range vsr {
			vs[i] = *vsr[i]
		}
	} else {
		for i := range vs {
			vs[i] = *vsr[i]
		}
	}
	return
}

func (vs VectorRefs) Cross(v Vector) {
	vs.ForEach((*Vector).Cross, v)
}

func (vs VectorRefs) Add(v Vector) {
	vs.ForEach((*Vector).Add, v)
}

func (vs VectorRefs) Subtract(v Vector) {
	vs.ForEach((*Vector).Subtract, v)
}

func (vs VectorRefs) Project(v Vector) {
	vs.ForEach((*Vector).Project, v)
}

func (vs VectorRefs) Product(m Matrix) {
	m.ForEachRef((*Vector).Product, vs)
}

func (vs VectorRefs) ProductT(m Matrix) {
	m.ForEachRef((*Vector).ProductT, vs)
}

func (vs VectorRefs) Sum() (v Vector) {
	v.AggregateRefs(vs, (*Vector).Add)
	return
}

func (vs VectorRefs) Multiply(s Scalar) {
	var multiply func(*Vector)
	multiply = func(v *Vector) {
		v.Multiply(s)
	}
	vs.ForEachNoParameter(multiply)
}

func (vs VectorRefs) Max() (v Vector) {
	v.Set(*vs[0])
	v.AggregateRefs(vs[1:], (*Vector).Max)
	return
}

func (vs VectorRefs) Min() (v Vector) {
	v.Set(*vs[0])
	v.AggregateRefs(vs[1:], (*Vector).Min)
	return
}

func (vs VectorRefs) Middle() (v Vector) {
	v = vs.Max()
	v.Mid(vs.Min())
	return
}

func (vs VectorRefs) Interpolate(v Vector, f Scalar) {
	f2 := 1 - f
	var interpolate func(*Vector, Vector)
	interpolate = func(v *Vector, v2 Vector) {
		v2.Multiply(f2)
		v.Multiply(f)
		v.Add(v2)
	}
	vs.ForEach(interpolate, v)
}

// apply a function repeatedly to the vector reference, parameterised by its current value and each vector in the supplied vectors in order.
func (v *Vector) AggregateRefs(vs VectorRefs, fn func(*Vector, Vector)) {
	for _, v2 := range vs {
		fn(v, *v2)
	}
}

func (vs VectorRefs) ForEach(fn func(*Vector, Vector), v Vector) {
	if !Parallel {
		vectorRefsApply(vs, fn, v)
	} else {
		vectorRefsApplyChunked(vs, fn, v)
	}
}

func vectorRefsApply(vs VectorRefs, fn func(*Vector, Vector), v Vector) {
	for i := range vs {
		fn(vs[i], v)
	}
}

func vectorRefsApplyChunked(vs VectorRefs, fn func(*Vector, Vector), v Vector) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range vectorRefsInChunks(vs, chunkSize(len(vs))) {
		running++
		go func(c VectorRefs) {
			vectorRefsApply(c, fn, v)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}

func (m Matrix) ForEachRef(fn func(*Vector, Matrix), vs VectorRefs) {
	if !Parallel {
		matrixApplyRef(vs, fn, m)
	} else {
		matrixApplyRefChunked(vs, fn, m)
	}
}

func matrixApplyRef(vs VectorRefs, fn func(*Vector, Matrix), m Matrix) {
	for i := range vs {
		fn(vs[i], m)
	}
}

func matrixApplyRefChunked(vs VectorRefs, fn func(*Vector, Matrix), m Matrix) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range vectorRefsInChunks(vs, chunkSize(len(vs))) {
		running++
		go func(c VectorRefs) {
			matrixApplyRef(c, fn, m)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}

// for each vector reference apply a function with no parameters
func (vs VectorRefs) ForEachNoParameter(fn func(*Vector)) {
	var inner func(*Vector, Vector)
	inner = func(v *Vector, _ Vector) {
		fn(v)
	}
	vs.ForEach(inner, Vector{})
}

func (vs Vectors) CrossAllRefs(vrs VectorRefs) {
	vs.ForAllRefs(vrs, (*Vector).Cross)
}

func (vs Vectors) AddAllRefs(vrs VectorRefs) {
	vs.ForAllRefs(vrs, (*Vector).Add)
}

func (vs Vectors) SubtractAllRefs(vrs VectorRefs) {
	vs.ForAllRefs(vrs, (*Vector).Subtract)
}

func (vs Vectors) ProjectAllRefs(vrs VectorRefs) {
	vs.ForAllRefs(vrs, (*Vector).Project)
}

func (vs Vectors) ForAllRefs(vrs VectorRefs, fn func(*Vector, Vector)) {
	if !Parallel {
		vectorsApplyAllRefs(vs, fn, vrs)
	} else {
		vectorsApplyAllRefsChunked(vs, fn, vrs)
	}
}

func vectorsApplyAllRefs(vs Vectors, fn func(*Vector, Vector), vs2 VectorRefs) {
	for i := range vs {
		fn(&vs[i], *vs2[i])
	}
}

func vectorsApplyAllRefsChunked(vs Vectors, fn func(*Vector, Vector), vrs VectorRefs) {
	done := make(chan struct{}, 1)
	var running uint
	// shorten vs to use only what we have in vs2rs
	if len(vs) > len(vrs) {
		vs = vs[:len(vrs)]
	}
	cs := chunkSize(len(vs))
	chunks2 := vectorRefsInChunks(vrs, cs)
	for chunk := range vectorsInChunks(vs, cs) {
		running++
		go func(c Vectors) {
			vectorsApplyAllRefs(c, fn, <-chunks2)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}

func (vrs VectorRefs) CrossAllRefs(vrs2 VectorRefs) {
	vrs.ForAllRefs(vrs2, (*Vector).Cross)
}

func (vrs VectorRefs) AddAllRefs(vrs2 VectorRefs) {
	vrs.ForAllRefs(vrs2, (*Vector).Add)
}

func (vrs VectorRefs) SubtractAllRefs(vrs2 VectorRefs) {
	vrs.ForAllRefs(vrs2, (*Vector).Subtract)
}

func (vrs VectorRefs) ProjectAllRefs(vrs2 VectorRefs) {
	vrs.ForAllRefs(vrs2, (*Vector).Project)
}

func (vrs VectorRefs) ForAllRefs(vrs2 VectorRefs, fn func(*Vector, Vector)) {
	if !Parallel {
		vectorRefsApplyAllRefs(vrs, fn, vrs2)
	} else {
		vectorRefsApplyAllChunkedRefs(vrs, fn, vrs2)
	}
}

func vectorRefsApplyAllRefs(vrs VectorRefs, fn func(*Vector, Vector), vrs2 VectorRefs) {
	for i := range vrs {
		fn(vrs[i], *vrs2[i])
	}
}

func vectorRefsApplyAllChunkedRefs(vrs VectorRefs, fn func(*Vector, Vector), vrs2 VectorRefs) {
	done := make(chan struct{}, 1)
	var running uint
	// shorten vs to use only what we have in vs2
	if len(vrs) > len(vrs2) {
		vrs = vrs[:len(vrs2)]
	}
	cs := chunkSize(len(vrs))
	chunks2 := vectorRefsInChunks(vrs2, cs)
	for chunk := range vectorRefsInChunks(vrs, cs) {
		running++
		go func(c VectorRefs) {
			vectorRefsApplyAllRefs(c, fn, <-chunks2)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}

func (vrs VectorRefs) CrossAll(vs Vectors) {
	vrs.ForAll(vs, (*Vector).Cross)
}

func (vrs VectorRefs) AddAll(vs Vectors) {
	vrs.ForAll(vs, (*Vector).Add)
}

func (vrs VectorRefs) SubtractAll(vs Vectors) {
	vrs.ForAll(vs, (*Vector).Subtract)
}

func (vrs VectorRefs) ProjectAll(vs Vectors) {
	vrs.ForAll(vs, (*Vector).Project)
}

func (vrs VectorRefs) ForAll(vs Vectors, fn func(*Vector, Vector)) {
	if !Parallel {
		vectorRefsApplyAll(vrs, fn, vs)
	} else {
		vectorRefsApplyAllChunked(vrs, fn, vs)
	}
}

func vectorRefsApplyAll(vs VectorRefs, fn func(*Vector, Vector), vs2 Vectors) {
	for i := range vs {
		fn(vs[i], vs2[i])
	}
}

func vectorRefsApplyAllChunked(vrs VectorRefs, fn func(*Vector, Vector), vs2 Vectors) {
	done := make(chan struct{}, 1)
	var running uint
	// shorten vrs to use only what we have in vs2
	if len(vrs) > len(vs2) {
		vrs = vrs[:len(vs2)]
	}
	cs := chunkSize(len(vrs))
	chunks2 := vectorsInChunks(vs2, cs)
	for chunk := range vectorRefsInChunks(vrs, cs) {
		running++
		go func(c VectorRefs) {
			vectorRefsApplyAll(c, fn, <-chunks2)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}

// return a sub selection of a VectorRefs with only references to Vector's that return true from the provided function.
func (vrs VectorRefs) Select(fn func(Vector) bool) (svs VectorRefs) {
	for _, vr := range vrs {
		if fn(*vr) {
			svs = append(svs, vr)
		}
	}
	return
}

// return a sub selection of a VectorRefs with only references at equal spaced strides.
func (vrs VectorRefs) Stride(s uint) (svs VectorRefs) {
	if s == 0 {
		return
	}
	is := int(s)
	svs = make(VectorRefs, len(vrs)/is+1)
	for i := range svs {
		svs[i] = vrs[i*is]
	}
	return
}

// return a slice of VectorRefs pointing to the Vector's from this that returned the slices index value from the provided function.
// or put another way;
// bin Vector's by the functions returned value.
// bins start at 1, a function returning a value of 0 causes the VecRef not to be in any of the returned bins.
func (vrs VectorRefs) Split(fn func(*Vector) uint) (ssvs []VectorRefs) {
	for _, vr := range vrs {
		ind := fn(vr)
		if ind > 0 {
			// pad, if needed, with a series of new VectorRefs to fill up to index. (max index not pre-known)
			if ind > uint(len(ssvs)) {
				ssvs = append(ssvs, make([]VectorRefs, ind-uint(len(ssvs)))...)
			}
			ssvs[ind-1] = append(ssvs[ind-1], vr)
		}
	}
	return
}

// for each vector apply a function with no parameters
func (vrs VectorRefs) ForEachInSlices(length, stride int, wrap bool, fn func(VectorRefs)) {
	if !Parallel {
		var i int
		for ; i < len(vrs)-length+1; i += stride {
			fn(vrs[i : i+length])
		}
		if wrap {
			joinSlice := make(VectorRefs, length, length)
			for ; i < len(vrs); i += stride {
				copy(joinSlice, vrs[i:])
				copy(joinSlice[len(vrs)-i:], vrs)
				fn(joinSlice)
			}
		}
	} else {
		vectorRefsInSlicesApplyChunked(vrs, length, stride, wrap, fn)
	}
}

func vectorRefsInSlicesApply(vrss []VectorRefs, fn func(VectorRefs)) {
	for _, vrs := range vrss {
		fn(vrs)
	}
}

func vectorRefsInSlicesApplyChunked(vrs VectorRefs, length, stride int, wrap bool, fn func(VectorRefs)) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range vectorRefsSlicesInChunks(vrs, chunkSize(len(vrs)), length, stride, wrap) {
		running++
		go func(c []VectorRefs) {
			vectorRefsInSlicesApply(c, fn)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}

// find the index in the Vectors that produces the lowest value from the function.
func (vrs VectorRefs) FindMin(toMin func(Vector) Scalar) int {
	if !Parallel || len(vrs) < chunkSize(len(vrs)) {
		return vectorRefsFindMin(vrs, toMin)
	} else {
		return vectorRefsFindMinChunked(vrs, toMin)
	}
}

func vectorRefsFindMin(vrs VectorRefs, toMin func(Vector) Scalar) (i int) {
	value := toMin(*vrs[0])
	var imv Scalar
	for j, jv := range vrs[1:] {
		imv = toMin(*jv)
		if imv < value {
			value, i = imv, j+1
		}
	}
	return
}

func vectorRefsFindMinChunked(vrs VectorRefs, toMin func(Vector) Scalar) (i int) {
	done := make(chan int, 1)
	var running uint
	for chunk := range vectorRefsInChunks(vrs, chunkSize(len(vrs))) {
		running++
		go func(c VectorRefs) {
			done <- vectorRefsFindMin(c, toMin)
		}(chunk)
	}
	//	if running==0 {return}
	i = <-done
	var j int
	for ; running > 1; running-- {
		j = <-done
		if toMin(*vrs[j]) < toMin(*vrs[i]) {
			i = j
		}
	}
	return
}

// search VectorRefs for the two Vector's that return the minimum value from the provided function
func (vrs VectorRefs) SearchMin(toMin func(Vector, Vector) Scalar) (i, j int) {
	j = vrs[1:].FindMin(func(v Vector) Scalar { return toMin(*vrs[0], v) }) + 1
	var jp int
	for ip, vip := range vrs[1 : len(vrs)-1] {
		jp = vrs[ip+2:].FindMin(func(v Vector) Scalar { return toMin(*vip, v) }) + 2 + ip
		if toMin(*vip, *vrs[jp]) < toMin(*vrs[i], *vrs[j]) {
			j, i = jp, ip+1
		}
	}
	return
}

//func (vrs VectorRefs) ApproxMiddle() Vector {
	// same nuber all sides?
//	for i:=0; i<len(vrs); i=+len(vrs)/20{
	 	
//	}
//}


// algorithm used only works when the function always returns a bigger value when points are further apart.
func (vrs VectorRefs) SearchMinRegionally(toMin func(Vector, Vector) Scalar) (iv, jv *Vector) {
	return vrs.SearchMinRegionallyCentered(vrs.Middle(), toMin)
}

// algorithm used only works when the function always returns a bigger value when points are further apart.
func (vrs VectorRefs) SearchMinRegionallyCentered(splitPoint Vector, toMin func(Vector, Vector) Scalar) (iv, jv *Vector) {
	// separate search into 8 regions, split by supplied points axes. then search regions along joins where the minimum  might have been missed due to crossing regions. 
	// if not more than 8 then could have situation where no region has a pair to search, so do non-split search on whole set, in this case.
	if len(vrs)<9{
		k, l := vrs.SearchMin(toMin) 
		iv,jv = vrs[k],vrs[l]
		return
	}
	axisSplitRegionChan:=vectorRefsSplitRegionally(vrs, splitPoint)
	var min Scalar
	// separate the first search to set min, which could be below zero.
	for vrss := range axisSplitRegionChan { 
		if len(vrss)>1{
			var kv,lv *Vector
			if len(vrss) > chunkSize(len(vrss)) {
				kv, lv = vrss.SearchMinRegionally(toMin) 
				}else{
				k, l := vrss.SearchMin(toMin) 
				kv,lv = vrss[k],vrss[l]
			}
			min=toMin(*kv,*lv)
			iv,jv=kv,lv
			break
		}
	}

	for vrss := range axisSplitRegionChan { 
		if len(vrss)>1{
			var kv,lv *Vector
			if len(vrss) > chunkSize(len(vrss)) {
				kv, lv = vrss.SearchMinRegionally(toMin) 
				}else{
				k, l := vrss.SearchMin(toMin) 
				kv,lv = vrss[k],vrss[l]
			}
			if 	klMin := toMin(*kv,*lv);klMin<min{
				min=klMin
				iv,jv=kv,lv
			}
		}
	}
	// also search 3 regions along joins
	// join regions width depends on the separation of min points already found
	dvx:=iv.x-jv.x
	if dvx<0 {dvx=-dvx}
	dvx+=splitPoint.x
	vrss:=vrs.Select(func(v Vector) bool {return v.x < dvx && v.x > -dvx})
	if len(vrss)>1 {
		k, l := vrss.SearchMin(toMin) 
		kv,lv := vrss[k],vrss[l]
		if 	klMin := toMin(*kv,*lv);klMin<min{
			min=klMin
			iv,jv=kv,lv 
		}
	}
	
	dvy:=iv.y-jv.y
	if dvy<0 {dvy=-dvy}
	dvy+=splitPoint.y
	vrss=vrs.Select(func(v Vector) bool {return v.y < dvy && v.y > -dvy})
	if len(vrss)>1 {
		k, l := vrss.SearchMin(toMin) 
		kv,lv := vrss[k],vrss[l]
		if 	klMin := toMin(*kv,*lv);klMin<min{
			min=klMin
			iv,jv=kv,lv
		}
	}	
	
	dvz:=iv.z-jv.z
	if dvz<0 {dvz=-dvz}
	dvz+=splitPoint.z
	vrss=vrs.Select(func(v Vector) bool {return v.z < dvz && v.z > -dvz})
	if len(vrss)>1 {
		k, l := vrss.SearchMin(toMin) 
		kv,lv := vrss[k],vrss[l]
		if 	klMin := toMin(*kv,*lv);klMin<min{
			min=klMin
			iv,jv=kv,lv
		}
	}
	return
}

