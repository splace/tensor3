package tensor3

type Matrices []Matrix

func (ms Matrices) Add(m Matrix) {
	ms.ForEach((*Matrix).Add, m)
}

func (ms Matrices) Subtract(m Matrix) {
	ms.ForEach((*Matrix).Subtract, m)
}

func (ms Matrices) Project(m Matrix) {
	ms.ForEach((*Matrix).Project, m)
}

func (ms Matrices) Product(m Matrix) {
	ms.ForEach((*Matrix).Product, m)
}

func (ms Matrices) ProductRight(m Matrix) {
	ms.ForEach((*Matrix).ProductRight, m)
}

func (ms Matrices) ProductT(m Matrix) {
	ms.ForEach((*Matrix).ProductT, m)
}

func (ms Matrices) ProductRightT(m Matrix) {
	ms.ForEach((*Matrix).ProductRightT, m)
}

func (ms Matrices) Transpose() {
	var transposer func(*Matrix, Matrix)
	transposer = func(m1 *Matrix, _ Matrix) {
		m1.Transpose()
	}
	ms.ForEach(transposer, Matrix{})
}

func (ms Matrices) Sum() (m Matrix) {
	m.Reduce(ms, (*Matrix).Add)
	return
}

func (ms Matrices) ReduceComponentWise(v Vector, fn func(*Vector, Vector)) {
	ms.ApplyComponentWiseVariac(v, fn, fn, fn)
}

func (ms Matrices) ApplyComponentWiseVariac(v Vector, fns ...interface{}) {
	done := make(chan struct{}, 1)
	var running uint
	switch len(fns) {
	case 3:
		if fn, ok := fns[2].(func(*Vector, Vector)); ok {
			running++
			go func() {
				vectorApply(ms, (*Matrix).applyZ, fn, v)
				done <- struct{}{}
			}()
		}
		fallthrough
	case 2:
		if fn, ok := fns[1].(func(*Vector, Vector)); ok {
			running++
			go func() {
				vectorApply(ms, (*Matrix).applyY, fn, v)
				done <- struct{}{}
			}()
		}
		fallthrough
	case 1:
		if fn, ok := fns[0].(func(*Vector, Vector)); ok {
			running++
			go func() {
				vectorApply(ms, (*Matrix).applyX, fn, v)
				done <- struct{}{}
			}()
		}
	}
	for ; running > 0; running-- {
		<-done
	}

}

func (ms Matrices) ForEach(fn func(*Matrix, Matrix), v Matrix) {
	if !Parallel {
		MatricesApply(ms, fn, v)
	} else {
		if Hints.ChunkSizeFixed {
			MatricesApplyChunked(ms, fn, v, Hints.DefaultChunkSize)
		} else {
			MatricesApplyChunked(ms, fn, v, Hints.DefaultChunkSize+uint(len(ms))/(Hints.CoresOverOne+1))
		}
	}
}

func MatricesApply(ms Matrices, fn func(*Matrix, Matrix), v Matrix) {
	for i := range ms {
		fn(&ms[i], v)
	}
}

func MatricesApplyChunked(ms Matrices, fn func(*Matrix, Matrix), v Matrix, chunkSize uint) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range MatricesInChunks(ms, chunkSize) {
		running++
		go func(c Matrices) {
			MatricesApply(c, fn, v)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}

func MatricesInChunks(ms Matrices, chunkSize uint) chan Matrices {
	c := make(chan Matrices)
	length := uint(len(ms))
	go func() {
		var bottom uint
		for top := chunkSize; top < length; top += chunkSize {
			c <- ms[bottom:top]
			bottom = top
		}
		c <- ms[bottom:]
		close(c)
	}()
	return c
}

func (ms Matrices) vectorApply(mfn func(*Matrix, func(*Vector, Vector), Vector), fn func(*Vector, Vector), v Vector) {
	if !Parallel {
		vectorApply(ms, mfn, fn, v)
	} else {
		if Hints.ChunkSizeFixed {
			vectorApplyChunked(ms, mfn, fn, v, Hints.DefaultChunkSize)
		} else {
			cs := uint(len(ms)) / (Hints.CoresOverOne + 1)
			if cs < Hints.DefaultChunkSize {
				cs = Hints.DefaultChunkSize
			}
			vectorApplyChunked(ms, mfn, fn, v, cs)
		}
	}
}

func vectorApply(ms Matrices, mfn func(*Matrix, func(*Vector, Vector), Vector), fn func(*Vector, Vector), v Vector) {
	for i := range ms {
		mfn(&ms[i], fn, v)
	}
}

func vectorApplyChunked(ms Matrices, mfn func(*Matrix, func(*Vector, Vector), Vector), fn func(*Vector, Vector), v Vector, chunkSize uint) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range MatricesInChunks(ms, chunkSize) {
		running++
		go func(c Matrices) {
			vectorApply(c, mfn, fn, v)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}