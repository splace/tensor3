package tensor3

type Vectors []Vector

func NewVectors(cs ...Scalar) (vs Vectors) {
	vs = make(Vectors, (len(cs)+2)/3)
	for i := range vs {
		vs[i] = NewVector(cs[i*3:]...)
	}
	return
}

func (vs Vectors) Index(vr *Vector) (i uint) {
	for j := range vs {
		if &vs[j] == vr {
			return uint(j + 1)
		}
	}
	return 0
}

func (vs Vectors) Cross(v Vector) {
	vs.ForEach((*Vector).Cross, v)
}

func (vs Vectors) Add(v Vector) {
	vs.ForEach((*Vector).Add, v)
}

func (vs Vectors) Subtract(v Vector) {
	vs.ForEach((*Vector).Subtract, v)
}

func (vs Vectors) Project(v Vector) {
	vs.ForEach((*Vector).Project, v)
}

func (vs Vectors) Sum() (v Vector) {
	v.ForAll(vs, (*Vector).Add)
	return
}

func (vs Vectors) Multiply(s Scalar) {
	var multiply func(*Vector)
	multiply = func(v *Vector) {
		v.Multiply(s)
	}
	vs.ForEachNoParameter(multiply)
}

func (vs Vectors) Max() (v Vector) {
	v.Set(vs[0])
	v.ForAll(vs[1:], (*Vector).Max)
	return
}

func (vs Vectors) Min() (v Vector) {
	v.Set(vs[0])
	v.ForAll(vs[1:], (*Vector).Min)
	return
}

func (vs Vectors) Middle() (v Vector) {
	v = vs.Max()
	v.Mid(vs.Min())
	return
}

func (vs Vectors) Interpolate(v Vector, f float64) {
	f1 := Base64(1 - f)
	f2 := Base64(f)
	var interpolate func(*Vector, Vector)
	interpolate = func(v *Vector, v2 Vector) {
		v2.Multiply(f1)
		v.Multiply(f2)
		v.Add(v2)
	}
	vs.ForEach(interpolate, v)
}

func (vs Vectors) ForEach(fn func(*Vector, Vector), v Vector) {
	if !Parallel {
		vectorsApply(vs, fn, v)
	} else {
		vectorsApplyChunked(vs, fn, v)
	}
}

func vectorsApply(vs Vectors, fn func(*Vector, Vector), v Vector) {
	for i := range vs {
		fn(&vs[i], v)
	}
}

func vectorsApplyChunked(vs Vectors, fn func(*Vector, Vector), v Vector) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range vectorsInChunks(vs, chunkSize(len(vs))) {
		running++
		go func(c Vectors) {
			vectorsApply(c, fn, v)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}

// for each vector apply a function with no parameters
func (vs Vectors) ForEachNoParameter(fn func(*Vector)) {
	var inner func(*Vector, Vector)
	inner = func(v *Vector, _ Vector) {
		fn(v)
	}
	vs.ForEach(inner, Vector{})
}

func (vs Vectors) CrossAll(vs2 Vectors) {
	vs.ForAll((*Vector).Cross, vs2)
}

func (vs Vectors) AddAll(vs2 Vectors) {
	vs.ForAll((*Vector).Add, vs2)
}

func (vs Vectors) SubtractAll(vs2 Vectors) {
	vs.ForAll((*Vector).Subtract, vs2)
}

func (vs Vectors) ProjectAll(vs2 Vectors) {
	vs.ForAll((*Vector).Project, vs2)
}

func (vs Vectors) ForAll(fn func(*Vector, Vector), vs2 Vectors) {
	if !Parallel {
		vectorsApplyAll(vs, fn, vs2)
	} else {
		vectorsApplyAllChunked(vs, fn, vs2)
	}
}

func vectorsApplyAll(vs Vectors, fn func(*Vector, Vector), vs2 Vectors) {
	for i := range vs {
		fn(&vs[i], vs2[i])
	}
}

func vectorsApplyAllChunked(vs Vectors, fn func(*Vector, Vector), vs2 Vectors) {
	done := make(chan struct{}, 1)
	var running uint
	// shorten vs to use only what we have in vs2
	if len(vs) > len(vs2) {
		vs = vs[:len(vs2)]
	}
	cs := chunkSize(len(vs))
	chunks2 := vectorsInChunks(vs2, cs)
	for chunk := range vectorsInChunks(vs, cs) {
		running++
		go func(c Vectors) {
			vectorsApplyAll(c, fn, <-chunks2)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}

// find the index in the Vectors that produces the lowest value from the provided function.
func (vs Vectors) FindMin(toMin func(Vector) Scalar) int {
	if !Parallel || len(vs) < chunkSize(len(vs)) {
		return vectorsFindMin(vs, toMin)
	} else {
		return vectorsFindMinChunked(vs, toMin)
	}
}

func vectorsFindMin(vs Vectors, toMin func(Vector) Scalar) (i int) {
	value := toMin(vs[0])
	var imv Scalar
	for j, jv := range vs[1:] {
		imv = toMin(jv)
		if imv < value {
			value, i = imv, j+1
		}
	}
	return
}

func vectorsFindMinChunked(vs Vectors, toMin func(Vector) Scalar) (i int) {
	done := make(chan int, 1)
	var running uint
	for chunk := range vectorsInChunks(vs, chunkSize(len(vs))) {
		running++
		go func(c Vectors) {
			done <- vectorsFindMin(c, toMin)
		}(chunk)
	}
	//	if running==0 {return}
	i = <-done
	var j int
	for ; running > 1; running-- {
		j = <-done
		if toMin(vs[j]) < toMin(vs[i]) {
			i = j
		}
	}
	return
}

// find the pair of indexes, from Vectors, that return the lowest value from the provided function.
func (vs Vectors) SearchMin(toMin func(Vector, Vector) Scalar) (i, j int) {
	j = vs[1:].FindMin(func(v Vector) Scalar { return toMin(vs[0], v) }) + 1
	var jp int
	for ip, vip := range vs[1 : len(vs)-1] {
		jp = vs[ip+2:].FindMin(func(v Vector) Scalar { return toMin(vip, v) }) + 2 + ip
		if toMin(vip, vs[jp]) < toMin(vs[i], vs[j]) {
			j, i = jp, ip+1
		}
	}
	return
	/*
		j=1
		value = toMin(vs[0], vs[1])
		var v1, v2 Vector
		var il, jl int
		for jl, v2 = range vs[2:] {
			nl := toMin(vs[0], v2)
			if nl < value {
				value, j = nl, jl+2
			}
		}

		for il, v1 = range vs[1:] {
			for jl, v2 = range vs[il+2:] {
				nl := toMin(v1, v2)
				if nl < value {
					value, i, j = nl, il+1, jl+il+2
				}
			}
		}
		return
	*/
}

func (vs Vectors) SearchMinRegionally(toMin func(Vector, Vector) Scalar) (i, j int) {
	return vs.SearchMinRegionallyCentered(vs.Middle(), toMin)
}

func (vs Vectors) SearchMinRegionallyCentered(splitPoint Vector, toMin func(Vector, Vector) Scalar) (i, j int) {
	//var notFirst bool
	for vrss := range vectorsSplitRegionally(vs, splitPoint) {
		if len(vrss) < 2 {
			continue
		} // rare, when used usefully, but single point still needs so be checked with 3 region edge points, and potentially needs perimeter crossing match test
		// TODO use go routine and channel for results.
		_, _ = vrss.SearchMin(toMin)

		// TODO region perimeter crossing match test, if either match point has lower toMin with any of its 3 region edge projected points, then another SearchMin, with the appropriate vrss, is needed.
		// TODO
	}
	return
}

/*
func (b *BaseType) Aggregate(vs Vectors,length,stride int, fn func(*BaseType, Vectors)) {
	for vsc := range vectorSlicesInChunks(vs,chunkSize(len(vs)),length,stride,false) {
		for _,vs:=range(vsc){
			fn(b, vs)
		}
	}
}
*/

// for each subsection of Vectors, obtained using the provided length and at the provided strides, apply a function.
// Note: when used in parallel, function application order isn't deterministic, so when used with overlapping subsections results are likewise non-deterministic.
func (vs Vectors) ForEachInSlices(length, stride int, wrap bool, fn func(Vectors)) {
	if !Parallel {
		var i int
		for ; i < len(vs)-length+1; i += stride {
			fn(vs[i : i+length])
		}
		if wrap {
			joinSlice := make(Vectors, length, length)
			for ; i < len(vs); i += stride {
				copy(joinSlice, vs[i:])
				copy(joinSlice[len(vs)-i:], vs)
				fn(joinSlice)
				copy(vs[i:], joinSlice)
				copy(vs, joinSlice[len(vs)-i:])
			}
		}
	} else {
		vectorsInSlicesApplyChunked(vs, length, stride, wrap, fn)
	}
}

func vectorsInSlicesApply(vss []Vectors, fn func(Vectors)) {
	for _, vs := range vss {
		fn(vs)
	}
}

func vectorsInSlicesApplyChunked(vs Vectors, length, stride int, wrap bool, fn func(Vectors)) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range vectorSlicesInChunks(vs, chunkSize(len(vs)), length, stride, wrap) {
		running++
		go func(c []Vectors) {
			vectorsInSlicesApply(c, fn)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}

// return a VectorRefs, with references from Vectors, that return true from the provided function.
func (vs Vectors) Select(fn func(*Vector) bool) (svs VectorRefs) {
	for i := range vs {
		if fn(&vs[i]) {
			svs = append(svs, &vs[i])
		}
	}
	return
}

// return a VectorRefs, with references from Vectors, that are at equal spaced strides.
func (vs Vectors) Stride(s uint) (svs VectorRefs) {
	if s == 0 {
		return
	}
	is := int(s)
	svs = make(VectorRefs, len(vs)/is+1)
	for i := range svs {
		svs[i] = &vs[i*is]
	}
	return
}

// return a slice of VectorRefs each referencing those Vector's that returned the slices index value from the provided function.
// or put another way;
// bin Vector's by the value returned by the provided function.
// bins start at 1, a function returning a value of 0 causes the VecRef not to be in any of the returned bins.
func (vs Vectors) Split(fn func(Vector) uint) (ssvs []VectorRefs) {
	for i := range vs {
		ind := fn(vs[i])
		if ind > 0 {
			// pad, if needed, with a series of new VectorRefs to fill up to index. (max index not preknown)
			if ind > uint(len(ssvs)) {
				ssvs = append(ssvs, make([]VectorRefs, ind-uint(len(ssvs)))...)
			}
			ssvs[ind-1] = append(ssvs[ind-1], &vs[i])
		}
	}
	return
}
