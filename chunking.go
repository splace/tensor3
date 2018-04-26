package tensor3

import "runtime"

func init() {
	Hints.Threads = uint(runtime.NumCPU()) - 1
	Hints.DefaultChunkSize = 10000
	//runtime.GOMAXPROCS(1)
}

var Hints struct {
	Threads          uint
	ChunkSizeFixed   bool
	DefaultChunkSize uint
}

// selects parallel application of functions to Vectors and Matrices types (slices of Vector and Matrix types).
// this occurs in chunks whose size is controlled by Hints.
var Parallel bool

// selects parallel application of functions to Matrix components,(its Vector fields).
// only improves performance if using costly functions, non of the built-ins are likely to benefit. YRMV.
var ParallelComponents bool



// return a channel of Vectors that are chunks of the passed Vectors
func vectorsInChunks(vs Vectors, chunkSize uint) chan Vectors {
	c := make(chan Vectors, 1)
	lastSplitMax := uint(len(vs))-chunkSize
	go func() {
		var bottom uint
		for top := chunkSize; top < lastSplitMax; top += chunkSize {
			c <- vs[bottom:top]
			bottom = top
		}
		c <- vs[bottom:]
		close(c)
	}()
	return c
}

// return a channel of VectorRefs that are chunks of the passed VectorRefs
func vectorRefsInChunks(vs VectorRefs, chunkSize uint) chan VectorRefs {
	c := make(chan VectorRefs, 1)
	lastSplitMax := uint(len(vs))-chunkSize
	go func() {
		var bottom uint
		for top := chunkSize; top < lastSplitMax; top += chunkSize {
			c <- vs[bottom:top]
			bottom = top
		}
		c <- vs[bottom:]
		close(c)
	}()
	return c
}

// return a channel of Matrices that are chunks of the passed Matrices
func matricesInChunks(ms Matrices, chunkSize uint) chan Matrices {
	c := make(chan Matrices)
	lastSplitMax := uint(len(ms))-chunkSize
	go func() {
		var bottom uint
		for top := chunkSize; top < lastSplitMax; top += chunkSize {
			c <- ms[bottom:top]
			bottom = top
		}
		c <- ms[bottom:]
		close(c)
	}()
	return c
}

