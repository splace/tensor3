package tensor3

type VectorRefs []*Vector

func NewVectorRefs(cs ...BaseType) (vs VectorRefs) {
	vs = make(VectorRefs, (len(cs)+2)/3)
	for i := range vs {
		v := NewVector(cs[i*3:]...)
		vs[i] = &v
	}
	return
}


// indexes are counting numbers, and in that the first vector has index 1.  
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

// rebases, maintaining values, a number of vectorrefs to point into a new returned vectors.
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

// convert memory Refs to array indexes, to retain information outside this execution context. 
// TODO find index from pointer, use unsafe? or read text, coule be much faster?
func (vsr VectorRefs) Indexes(vs Vectors) (is []uint) {
	is = make([]uint, len(vsr))
	for ir, r := range vsr {
		for i := range vs {
			if &vs[i] == r {
				is[ir] = uint(i + 1)
				break
			}
		}
	}
	return is
}

// make a slice of Vectors from a slice of Vector references. 
func (vsr VectorRefs) Dereference() (vs Vectors) {
	vs = make(Vectors, len(vsr))
	for i := range vs {
		vs[i] = *vsr[i]
	}
	return
}

// make a slice of VectorRefs from a slice of Vectors. 
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
	v.ForAllRefs(vs, (*Vector).Add)
	return
}

func (vs VectorRefs) Multiply(s BaseType) {
	var multiply func(*Vector)
	multiply = func(v *Vector) {
		v.Multiply(s)
	}
	vs.ForEachNoParameter(multiply)
}

func (vs VectorRefs) Max() (v Vector) {
	v.Set(*vs[0])
	v.ForAllRefs(vs[1:], (*Vector).Max)
	return
}

func (vs VectorRefs) Min() (v Vector) {
	v.Set(*vs[0])
	v.ForAllRefs(vs[1:], (*Vector).Min)
	return
}

func (vs VectorRefs) Interpolate(v Vector, f BaseType) {
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
	for chunk := range vectorRefsInChunks(vs) {
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
	for chunk := range vectorRefsInChunks(vs) {
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

// apply a function without a vector parameter using a dummy
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
	chunks2 := vectorRefsInChunks(vrs)
	for chunk := range vectorsInChunks(vs) {
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
	chunks2 := vectorRefsInChunks(vrs2)
	for chunk := range vectorRefsInChunks(vrs) {
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

func vectorRefsApplyAllChunked(vs VectorRefs, fn func(*Vector, Vector), vs2 Vectors) {
	done := make(chan struct{}, 1)
	var running uint
	chunks2 := vectorsInChunks(vs2)
	for chunk := range vectorRefsInChunks(vs) {
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



