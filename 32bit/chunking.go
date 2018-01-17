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
