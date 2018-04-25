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
	v.AggregateRefs(vs, (*Vector).Add)
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
	v.AggregateRefs(vs[1:], (*Vector).Max)
	return
}

func (vs VectorRefs) Min() (v Vector) {
	v.Set(*vs[0])
	v.AggregateRefs(vs[1:], (*Vector).Min)
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

func (vs VectorRefs) ForEach(fn func(*Vector, Vector), v Vector) {
	if !Parallel {
		vectorRefsApply(vs, fn, v)
	} else {
		if Hints.ChunkSizeFixed {
			vectorRefsApplyChunked(vs, fn, v, Hints.DefaultChunkSize)
		} else {
			cs := uint(len(vs)) / (Hints.Threads + 1)
			if cs < Hints.DefaultChunkSize {
				cs = Hints.DefaultChunkSize
			}
			vectorRefsApplyChunked(vs, fn, v, cs)
		}
	}
}

func vectorRefsApply(vs VectorRefs, fn func(*Vector, Vector), v Vector) {
	for i := range vs {
		fn(vs[i], v)
	}
}

func vectorRefsApplyChunked(vs VectorRefs, fn func(*Vector, Vector), v Vector, chunkSize uint) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range vectorRefsInChunks(vs, chunkSize) {
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

func vectorRefsInChunks(vs VectorRefs, chunkSize uint) chan VectorRefs {
	c := make(chan VectorRefs, 1)
	length := uint(len(vs))
	go func() {
		var bottom uint
		for top := chunkSize; top < length; top += chunkSize {
			c <- vs[bottom:top]
			bottom = top
		}
		c <- vs[bottom:]
		close(c)
	}()
	return c
}

func (m Matrix) ForEachRef(fn func(*Vector, Matrix), vs VectorRefs) {
	if !Parallel {
		matrixApplyRef(vs, fn, m)
	} else {
		if Hints.ChunkSizeFixed {
			matrixApplyRefChunked(vs, fn, m, Hints.DefaultChunkSize)
		} else {
			cs := uint(len(vs)) / (Hints.Threads + 1)
			if cs < Hints.DefaultChunkSize {
				cs = Hints.DefaultChunkSize
			}
			matrixApplyRefChunked(vs, fn, m, cs)
		}
	}
}

func matrixApplyRef(vs VectorRefs, fn func(*Vector, Matrix), m Matrix) {
	for i := range vs {
		fn(vs[i], m)
	}
}

func matrixApplyRefChunked(vs VectorRefs, fn func(*Vector, Matrix), m Matrix, chunkSize uint) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range vectorRefsInChunks(vs, chunkSize) {
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
		if Hints.ChunkSizeFixed {
			vectorsApplyAllRefsChunked(vs, fn, vrs, Hints.DefaultChunkSize)
		} else {
			cs := uint(len(vs)) / (Hints.Threads + 1)
			if cs < Hints.DefaultChunkSize {
				cs = Hints.DefaultChunkSize
			}
			vectorsApplyAllRefsChunked(vs, fn, vrs, cs)
		}
	}
}

func vectorsApplyAllRefs(vs Vectors, fn func(*Vector, Vector), vs2 VectorRefs) {
	for i := range vs {
		fn(&vs[i], *vs2[i])
	}
}

func vectorsApplyAllRefsChunked(vs Vectors, fn func(*Vector, Vector), vrs VectorRefs, chunkSize uint) {
	done := make(chan struct{}, 1)
	var running uint
	chunks2 := vectorRefsInChunks(vrs, chunkSize)
	for chunk := range vectorsInChunks(vs, chunkSize) {
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
		if Hints.ChunkSizeFixed {
			vectorRefsApplyAllChunkedRefs(vrs, fn, vrs2, Hints.DefaultChunkSize)
		} else {
			cs := uint(len(vrs)) / (Hints.Threads + 1)
			if cs < Hints.DefaultChunkSize {
				cs = Hints.DefaultChunkSize
			}
			vectorRefsApplyAllChunkedRefs(vrs, fn, vrs2, cs)
		}
	}
}

func vectorRefsApplyAllRefs(vrs VectorRefs, fn func(*Vector, Vector), vrs2 VectorRefs) {
	for i := range vrs {
		fn(vrs[i], *vrs2[i])
	}
}

func vectorRefsApplyAllChunkedRefs(vrs VectorRefs, fn func(*Vector, Vector), vrs2 VectorRefs, chunkSize uint) {
	done := make(chan struct{}, 1)
	var running uint
	chunks2 := vectorRefsInChunks(vrs2, chunkSize)
	for chunk := range vectorRefsInChunks(vrs, chunkSize) {
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
		if Hints.ChunkSizeFixed {
			vectorRefsApplyAllChunked(vrs, fn, vs, Hints.DefaultChunkSize)
		} else {
			cs := uint(len(vs)) / (Hints.Threads + 1)
			if cs < Hints.DefaultChunkSize {
				cs = Hints.DefaultChunkSize
			}
			vectorRefsApplyAllChunked(vrs, fn, vs, cs)
		}
	}
}

func vectorRefsApplyAll(vs VectorRefs, fn func(*Vector, Vector), vs2 Vectors) {
	for i := range vs {
		fn(vs[i], vs2[i])
	}
}

func vectorRefsApplyAllChunked(vs VectorRefs, fn func(*Vector, Vector), vs2 Vectors, chunkSize uint) {
	done := make(chan struct{}, 1)
	var running uint
	chunks2 := vectorsInChunks(vs2, chunkSize)
	for chunk := range vectorRefsInChunks(vs, chunkSize) {
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



