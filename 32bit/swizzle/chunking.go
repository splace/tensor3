package tensor3

import "runtime"

func init() {
	Hints.CoresOverOne = uint(runtime.NumCPU()) - 1
	Hints.DefaultChunkSize = 10000
	//runtime.GOMAXPROCS(1)
}

var Hints struct {
	CoresOverOne     uint
	ChunkSizeFixed   bool
	DefaultChunkSize uint
}

var Parallel bool
