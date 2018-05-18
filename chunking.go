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

func chunkSize(l uint) uint {
	if !Hints.ChunkSizeFixed {
		if cs := l / (Hints.Threads + 1); cs > Hints.DefaultChunkSize {
			return cs
		}
	}
	return Hints.DefaultChunkSize
}

// return a channel of Vectors that are chunks of the passed Vectors
func vectorsInChunks(vs Vectors) chan Vectors {
	c := make(chan Vectors, 1)
	cs:=chunkSize(uint(len(vs)))
	lastSplitMax := uint(len(vs))-cs
	go func() {
		var bottom uint
		for top := cs; top < lastSplitMax; top += cs {
			c <- vs[bottom:top]
			bottom = top
		}
		c <- vs[bottom:]
		close(c)
	}()
	return c
}

// return a channel of Matrices that are chunks of the passed Matrices
func matricesInChunks(ms Matrices) chan Matrices {
	c := make(chan Matrices)
	cs:=chunkSize(uint(len(ms)))
	lastSplitMax := uint(len(ms))-cs
	go func() {
		var bottom uint
		for top := cs; top < lastSplitMax; top += cs {
			c <- ms[bottom:top]
			bottom = top
		}
		c <- ms[bottom:]
		close(c)
	}()
	return c
}

// return a channel of VectorRefs that are chunks of the passed VectorRefs
func vectorRefsInChunks(vs VectorRefs) chan VectorRefs {
	c := make(chan VectorRefs, 1)
	cs:=chunkSize(uint(len(vs)))
	lastSplitMax := uint(len(vs))-cs
	go func() {
		var bottom uint
		for top := cs; top < lastSplitMax; top += cs {
			c <- vs[bottom:top]
			bottom = top
		}
		c <- vs[bottom:]
		close(c)
	}()
	return c
}


// return a channel of VectorRefs that are chunks of the passed VectorRefs.
// as an optimisation, which some functions might benefit from, the VectorRefs are reordered so that each chunk contains all/only the values within a region. nearby points are MUCH more likely to be in the same chunk.
func vectorRefsInRegionalChunks(vs VectorRefs) chan VectorRefs {
	c := make(chan VectorRefs, 1)
	cs:=chunkSize(uint(len(vs)))
	if cs>uint(len(vs)){
		c <- vs
		close(c)
		return c
	}
	go func() {
		// sample 5% of points to make guess at distribution
		// TODO sample cubed root number of points?
		sp:=make(VectorRefs,len(vs)/20)
		for i:=range(sp){
			sp[i]=vs[i*20]
		}
		// TODO improve this fixed 8-way only scheme. but if less than 8 cores, is there much point?
		// TODO regions to give chunks of a size similar to other chunking functions. 
		average:=sp.Sum()
		divisor:=BaseType(len(sp))
		average.x/=divisor
		average.y/=divisor
		average.z/=divisor
		//(&average).Divide(BaseType(len(sp)))
		var chunks [2][2][2]VectorRefs
		for _,v := range(vs){
			if v.x>average.x {
				if v.y>average.y {
					if v.z>average.z {
						chunks[1][1][1]=append(chunks[1][1][1],v)
					}else{
						chunks[1][1][0]=append(chunks[1][1][0],v)
					}
				}else{
					if v.z>average.z {
						chunks[1][0][1]=append(chunks[1][0][1],v)
					}else{
						chunks[1][0][0]=append(chunks[1][0][0],v)
					}
				}
			}else{
				if v.y>average.y {
					if v.z>average.z {
						chunks[0][1][1]=append(chunks[0][1][1],v)
					}else{
						chunks[0][1][0]=append(chunks[0][1][0],v)
					}
				}else{
					if v.z>average.z {
						chunks[0][0][1]=append(chunks[0][0][1],v)
					}else{
						chunks[0][0][0]=append(chunks[0][0][0],v)
					}
				}
			}
		}
		c <- chunks[0][0][0]
		c <- chunks[0][0][1]
		c <- chunks[0][1][0]
		c <- chunks[0][1][1]
		c <- chunks[1][0][0]
		c <- chunks[1][0][1]
		c <- chunks[1][1][0]
		c <- chunks[1][1][1]
		close(c)
	}()
	return c
}



