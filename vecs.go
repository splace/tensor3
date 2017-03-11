package tensor3

type Vectors []Vector

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

func (vs Vectors) Product(m Matrix) {
	m.ForEach((*Vector).Product, vs)
}

func (vs Vectors) ProductT(m Matrix) {
	m.ForEach((*Vector).ProductT, vs)
}

func (vs Vectors) Sum() (v Vector) {
	v.Reduce(vs, (*Vector).Add)
	return
}

func (vs Vectors) Max() (v Vector) {
	v.Set(vs[0])
	v.Reduce(vs[1:], (*Vector).Max)
	return
}

func (vs Vectors) Min() (v Vector) {
	v.Set(vs[0])
	v.Reduce(vs[1:], (*Vector).Min)
	return
}

func (vs Vectors) Interpolate(v Vector, f BaseType) {
	f2 := 1 - f
	var interpolate func(*Vector, Vector)
	interpolate = func(v *Vector, v2 Vector) {
		v2.Multiply(f2)
		v.Multiply(f)
		v.Add(v2)
	}
	vs.ForEach(interpolate, v)
}

func (vs Vectors) ForEach(fn func(*Vector, Vector), v Vector) {
	if !Parallel {
		vectorsApply(vs, fn, v)
	} else {
		if Hints.ChunkSizeFixed {
			vectorsApplyChunked(vs, fn, v, Hints.DefaultChunkSize)
		} else {
			cs := uint(len(vs)) / (Hints.Threads + 1)
			if cs < Hints.DefaultChunkSize {
				cs = Hints.DefaultChunkSize
			}
			vectorsApplyChunked(vs, fn, v, cs)
		}
	}
}

func vectorsApply(vs Vectors, fn func(*Vector, Vector), v Vector) {
	for i := range vs {
		fn(&vs[i], v)
	}
}

func vectorsApplyChunked(vs Vectors, fn func(*Vector, Vector), v Vector, chunkSize uint) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range vectorsInChunks(vs, chunkSize) {
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

func vectorsInChunks(vs Vectors, chunkSize uint) chan Vectors {
	c := make(chan Vectors, 1)
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

func (m Matrix) ForEach(fn func(*Vector, Matrix), vs Vectors) {
	if !Parallel {
		matrixApply(vs, fn, m)
	} else {
		if Hints.ChunkSizeFixed {
			matrixApplyChunked(vs, fn, m, Hints.DefaultChunkSize)
		} else {
			cs := uint(len(vs)) / (Hints.Threads + 1)
			if cs < Hints.DefaultChunkSize {
				cs = Hints.DefaultChunkSize
			}
			matrixApplyChunked(vs, fn, m, cs)
		}
	}
}

func matrixApply(vs Vectors, fn func(*Vector, Matrix), m Matrix) {
	for i := range vs {
		fn(&vs[i], m)
	}
}

func matrixApplyChunked(vs Vectors, fn func(*Vector, Matrix), m Matrix, chunkSize uint) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range vectorsInChunks(vs, chunkSize) {
		running++
		go func(c Vectors) {
			matrixApply(c, fn, m)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}
