package tensor3

import "runtime"

func init() {
	Hints.Threads = runtime.NumCPU() - 1
	Hints.DefaultChunkSize = 10000
}

var Hints struct {
	Threads          int   // used to stop unnecessarily making more chunks than available cores.
	ChunkSizeFixed   bool
	DefaultChunkSize int   // ideally set at run time to number of items that fit into cpu cache.
}

// selects parallel application of functions to Vectors and Matrices types (slices of Vector and Matrix types).
// this occurs in chunks whose size is controlled by Hints.
var Parallel bool

// selects parallel application of functions to Matrix components,(its Vector fields).
// only improves performance if using costly functions, non of the built-ins are likely to benefit. YRMV.
var ParallelComponents bool

// use hints to calc a good chunk size
func chunkSize(l int) int {
	if !Hints.ChunkSizeFixed {
		if cs := l / Hints.Threads; cs > Hints.DefaultChunkSize {
			return cs
		}
	}
	return Hints.DefaultChunkSize
}

// return a channel of Vectors that are chunks of the passed Vectors
func vectorsInChunks(vs Vectors, cs int) chan Vectors {
	c := make(chan Vectors, 1)
	lastSplitMax := len(vs)-cs/2
	go func() {
		var bottom int
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
func matricesInChunks(ms Matrices, cs int) chan Matrices {
	c := make(chan Matrices)
	lastSplitMax := len(ms)-cs/2
	go func() {
		var bottom int
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
func vectorRefsInChunks(vs VectorRefs, cs int) chan VectorRefs {
	c := make(chan VectorRefs, 1)
	lastSplitMax := len(vs)-cs/2
	go func() {
		var bottom int
		for top := cs; top < lastSplitMax; top += cs {
			c <- vs[bottom:top]
			bottom = top
		}
		c <- vs[bottom:]
		close(c)
	}()
	return c
}


// return a channel of chunks of, fixed length slices of, the passed Vectors
// progress by Stride Vector's for each slice, if Stride less than length, the same Vector can appear in consecutive slices.
// if wrap true, include slices that wrap around, from the end to the start of the passed Vectors.
// notice: all the slices will be the same provided length.
// notice: can panic if length larger than chunksize/2 (ie the min. terminal chunk size)  
func vectorSlicesInChunks(vs Vectors, cs int,length,stride int, wrap bool) chan []Vectors {
	c := make(chan []Vectors, 2)  // 2 so that the next chunk is being calculated in parallel, here unlike other chunking it has a significant cost, although if all cores kept 100% busy, not beneficial.
	go func(){
		// need to special case last chunk; it might have to include the wrap-round's, but don't know its the last until channel closes, so handle previous loop cycle.
		chunkChan :=vectorsInChunks(vs,cs)
		firstChunk := <- chunkChan // keep first chunk for potential wrap-round
		previousChunk := firstChunk
		var i int
		for chunk:=range chunkChan{
			vssc := make([]Vectors,len(previousChunk)/stride)
			for ;i<len(vssc);i+=stride {
				vssc[i]=previousChunk[i:i+length]
			}
			c <- vssc
			previousChunk=chunk
			i%=stride  // reset start index, allowing for stride continuation
		}
		// now handle the last (previous) chunk
		if wrap {
			vssc := make([]Vectors,len(previousChunk)/stride)
			for ; i< len(vssc)-length+1;i+=stride {
				vssc[i]=previousChunk[i:i+length]
			}
			// add the overlapping slices
			for ;i < len(vssc);i+=stride {
				vssc[i]=previousChunk[i:]
				vssc[i]=append(vssc[i],firstChunk[:length-len(vssc[i])]...)
			}			
			c <- vssc
		}else{
			// not wrapping so its just shortened
			vssc := make([]Vectors,(len(previousChunk)-length+1)/stride)
			for ;i < len(vssc);i+=stride {
				vssc[i]=previousChunk[i:i+length]
			}
			c <- vssc
		}
		close(c)
	}()
	return c
}

// TODO VectorRefSlicesInChunks(vs Vectors, cs int,length,stride int, wrap bool) chan []Vectors {


// return a channel of VectorRefs that are chunks of the passed VectorRefs.
// as an optimisation, which some functions might benefit from, the VectorRefs are reordered so that each chunk contains all/only the values within a spacial region, so nearby points are MUCH more likely to be in the same chunk.
// re-apply on returned chunks, recursively,  to sub-divide
func vectorRefsInRegionalChunks(vs VectorRefs, centre Vector, cs int) chan VectorRefs {
	// TODO continue to subdivide if exceed chunk size 
	// TODO return, another channel?, boundingbox of chunk? 
	c := make(chan VectorRefs, 8) 
	if cs>len(vs){
		c <- vs
		close(c)
		return c
	}
	go func() {
		// sample 5% of points to make guess at distribution
		sp:=make(VectorRefs,len(vs)/20)
		for i:=range(sp){
			sp[i]=vs[i*20]
		}
		var chunks [2][2][2]VectorRefs
		for _,v := range(vs){
			if v.x>centre.x {
				if v.y>centre.y {
					if v.z>centre.z {
						chunks[1][1][1]=append(chunks[1][1][1],v)
					}else{
						chunks[1][1][0]=append(chunks[1][1][0],v)
					}
				}else{
					if v.z>centre.z {
						chunks[1][0][1]=append(chunks[1][0][1],v)
					}else{
						chunks[1][0][0]=append(chunks[1][0][0],v)
					}
				}
			}else{
				if v.y>centre.y {
					if v.z>centre.z {
						chunks[0][1][1]=append(chunks[0][1][1],v)
					}else{
						chunks[0][1][0]=append(chunks[0][1][0],v)
					}
				}else{
					if v.z>centre.z {
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



