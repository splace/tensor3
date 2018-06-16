package tensor3

import "runtime"

func init() {
	Hints.Threads = runtime.NumCPU() - 1
	Hints.DefaultChunkSize = 10000
}

// parallel optimisation Hints.
var Hints struct {
	Threads          int   // used to keep chunk to core ratio near parity.
	ChunkSizeFixed   bool
	DefaultChunkSize int   // ideally set at run time to number of items that fit into CPU cache.
}

// selects parallel application of functions to Vectors and Matrices types (slices of Vector and Matrix types).
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


// return a channel of slices from the passed Vectors.
// the inner slices are al the same, provided, Length.
// the start of each slice is spaced by Stride.
// if wrap is true then the last Vector is considered to join to the first.
// notice: wrapped around slices are newly created, modifying their content, unlike non-wrapped, won't change the source Vectors, if consistent behaviour is needed use VectorRefs chunks.
// notice: if the Stride is less than the Length, then the same Vector will appear in consecutive slices.
// notice: can panic if Length larger than chunksize/2 (ie the min. possible size of the last chunk)  
func vectorSlicesInChunks(vs Vectors, cs,length,stride int, wrap bool) chan []Vectors {
	c := make(chan []Vectors, 2)  // 2 so that the next chunk is being calculated in parallel, here unlike other chunking it has a significant cost, although if all cores kept 100% busy, not beneficial.
	go func(){
		// need to special case last chunk; it might have to include the wrap-round's, but don't know its the last until channel closes, so handle previous loop cycle.
		chunkChan :=vectorsInChunks(vs,cs)
		firstChunk := <- chunkChan // keep first chunk for potential wrap-round
		previousChunk := firstChunk
		var i int
		// TODO all array lengths could be precalculate instead of using append
		for chunk:=range chunkChan{
			var vssc []Vectors
			for ;i<len(previousChunk);i+=stride {
				vssc=append(vssc,previousChunk[i:i+length])
			}
			c <- vssc
			i%=len(previousChunk)  // reset start index, allowing for stride continuation across chunks
			previousChunk=chunk
		}
		// now handle the last (previous) chunk
		if wrap {
			var vssc []Vectors
			for ;i<len(previousChunk)-length+1;i+=stride {
				vssc=append(vssc,previousChunk[i:i+length])
			}
			// add the beginning Vectors
			for ;i<len(previousChunk);i+=stride {
				vssc=append(vssc,previousChunk[i:])
				vssc[len(vssc)-1]=append(vssc[len(vssc)-1],firstChunk[:length-len(vssc[len(vssc)-1])]...)
			}			
			c <- vssc
		}else{
			// not wrapping so its just shortened
			var vssc []Vectors
			for ;i<len(previousChunk)-length+1;i+=stride {
				vssc=append(vssc,previousChunk[i:i+length])
			}
			c <- vssc
		}
		close(c)
	}()
	return c
}


// see vectorSlicesInChunks
func vectorRefsSlicesInChunks(vs VectorRefs, cs,length,stride int, wrap bool) chan []VectorRefs {
	c := make(chan []VectorRefs, 2)  // 2 so that the next chunk is being calculated in parallel, here unlike other chunking it has a significant cost, although if all cores kept 100% busy, not beneficial.
	go func(){
		// need to special case last chunk; it might have to include the wrap-round's, but don't know its the last until channel closes, so handle previous loop cycle.
		chunkChan :=vectorRefsInChunks(vs,cs)
		firstChunk := <- chunkChan // keep first chunk for potential wrap-round
		previousChunk := firstChunk
		var i int
		// TODO all array lengths could be precalculate instead of using append
		for chunk:=range chunkChan{
			var vssc []VectorRefs
			for ;i<len(previousChunk);i+=stride {
				vssc=append(vssc,previousChunk[i:i+length])
			}
			c <- vssc
			i%=len(previousChunk)  // reset start index, allowing for stride continuation across chunks
			previousChunk=chunk
		}
		// now handle the last (previous) chunk
		if wrap {
			var vssc []VectorRefs
			for ;i<len(previousChunk)-length+1;i+=stride {
				vssc=append(vssc,previousChunk[i:i+length])
			}
			// add the beginning Vectors
			for ;i<len(previousChunk);i+=stride {
				vssc=append(vssc,previousChunk[i:])
				vssc[len(vssc)-1]=append(vssc[len(vssc)-1],firstChunk[:length-len(vssc[len(vssc)-1])]...)
			}			
			c <- vssc
		}else{
			// not wrapping so its just shortened
			var vssc []VectorRefs
			for ;i<len(previousChunk)-length+1;i+=stride {
				vssc=append(vssc,previousChunk[i:i+length])
			}
			c <- vssc
		}
		close(c)
	}()
	return c
}

// see vectorSlicesInChunks
func vectorRefsInMatrixChunks(vs VectorRefs, cs,stride int, wrap bool) chan Matrices {
	c := make(chan Matrices, 2)  // 2 so that the next chunk is being calculated in parallel, here unlike other chunking it has a significant cost, although if all cores kept 100% busy, not beneficial.
	go func(){
		// need to special case last chunk; it might have to include the wrap-round's, but don't know its the last until channel closes, so handle previous loop cycle.
		chunkChan :=vectorRefsInChunks(vs,cs)
		firstChunk := <- chunkChan // keep first chunk for potential wrap-round
		previousChunk := firstChunk
		var i int
		// TODO all array lengths could be precalculate instead of using append
		for chunk:=range chunkChan{
			var mssc Matrices
			for ;i<len(previousChunk);i+=stride {
				mssc=append(mssc,Matrix{*previousChunk[i],*previousChunk[i+1],*previousChunk[i+2]})
			}
			c <- mssc
			i%=len(previousChunk)  // reset start index, allowing for stride continuation across chunks
			previousChunk=chunk
		}
		// now handle the last (previous) chunk
		if wrap {
			var mssc Matrices
			for ;i<len(previousChunk)-2;i+=stride {
				mssc=append(mssc,Matrix{*previousChunk[i],*previousChunk[i+1],*previousChunk[i+2]})
			}
			// add the beginning Vectors
			for ;i<len(previousChunk);i+=stride {
				if i==len(previousChunk)-1{
					mssc=append(mssc,Matrix{*previousChunk[i],*firstChunk[1],*firstChunk[2]})
				}else{
					mssc=append(mssc,Matrix{*previousChunk[i-1],*firstChunk[i],*firstChunk[1]})
				}
			}			
			c <- mssc
		}else{
			// not wrapping so its just shortened
			var mssc Matrices
			for ;i<len(previousChunk)-2;i+=stride {
				mssc=append(mssc,Matrix{*previousChunk[i],*previousChunk[i+1],*previousChunk[i+2]})
			}
			c <- mssc
		}
		close(c)
	}()
	return c
}

// see vectorSlicesInChunks
func vectorsInMatrixChunks(vs Vectors, cs,stride int, wrap bool) chan Matrices {
	c := make(chan Matrices, 2)  // 2 so that the next chunk is being calculated in parallel, here unlike other chunking it has a significant cost, although if all cores kept 100% busy, not beneficial.
	go func(){
		// need to special case last chunk; it might have to include the wrap-round's, but don't know its the last until channel closes, so handle previous loop cycle.
		chunkChan :=vectorsInChunks(vs,cs)
		firstChunk := <- chunkChan // keep first chunk for potential wrap-round
		previousChunk := firstChunk
		var i int
		// TODO all array lengths could be precalculate instead of using append
		for chunk:=range chunkChan{
			var mssc Matrices
			for ;i<len(previousChunk);i+=stride {
				mssc=append(mssc,Matrix{previousChunk[i],previousChunk[i+1],previousChunk[i+2]})
			}
			c <- mssc
			i%=len(previousChunk)  // reset start index, allowing for stride continuation across chunks
			previousChunk=chunk
		}
		// now handle the last (previous) chunk
		if wrap {
			var mssc Matrices
			for ;i<len(previousChunk)-2;i+=stride {
				mssc=append(mssc,Matrix{previousChunk[i],previousChunk[i+1],previousChunk[i+2]})
			}
			// add the beginning Vectors
			for ;i<len(previousChunk);i+=stride {
				if i==len(previousChunk)-1{
					mssc=append(mssc,Matrix{previousChunk[i],firstChunk[1],firstChunk[2]})
				}else{
					mssc=append(mssc,Matrix{previousChunk[i-1],firstChunk[i],firstChunk[1]})
				}
			}			
			c <- mssc
		}else{
			// not wrapping so its just shortened
			var mssc Matrices
			for ;i<len(previousChunk)-2;i+=stride {
				mssc=append(mssc,Matrix{previousChunk[i],previousChunk[i+1],previousChunk[i+2]})
			}
			c <- mssc
		}
		close(c)
	}()
	return c
}

// return a channel of VectorRefs that are chunks of the passed VectorRefs.
// as an optimisation, which some functions might benefit from, the VectorRefs are split so that each chunk contains all/only the VectorRefs within a spacial region, meaning nearby points are MUCH more likely to be in the same chunk.
func vectorRefsSplitRegionally(vrs VectorRefs, centre Vector) chan VectorRefs {
	// TODO continue to subdivide if exceed chunk size?
	// TODO return, another channel?, boundingbox of chunk? 
	cvr := make(chan VectorRefs,8)  // non blocking since a max on 8 VectorRefs from this split function
	// range over a slice of VectorRefs, returned by the Split function, using a function that splits into 8 regions using which side of the origin Vector, by axis alignment, the point is on.
	for _,s:=range func() []VectorRefs {
		return vrs.Split(
			func(v *Vector)(i uint){
				i++   // index never zero, since for this all points go somewhere
				if v.x>=centre.x {i++}
				if v.y>=centre.y {i+=2} 
				if v.z>=centre.z {i+=4}
				return
			},
		)
	}(){
		cvr <- s
	}
	close(cvr)
	return cvr
}

// return a channel of VectorRefs that are chunks of the passed Vectors.
// as an optimisation, which some functions might benefit from, the Vectors are split so that each chunk contains all/only the VectorRefs within a spacial region, meaning nearby points are MUCH more likely to be in the same chunk.
func vectorsSplitRegionally(vs Vectors, centre Vector) chan VectorRefs {
	// TODO continue to subdivide if exceed chunk size?
	// TODO return, another channel?, boundingbox of chunk? 
	cvr := make(chan VectorRefs,8)  // non blocking since a max on 8 VectorRefs from this split function
	// range over a slice of VectorRefs, returned by the Split function, using a function that splits into 8 regions using which side of the origin Vector, by axis alignment, the point is on.
	for _,s:=range func() []VectorRefs {
		return vs.Split(
			func(v Vector)(i uint){
				i++   // index never zero, since for this all points go somewhere
				if v.x>=centre.x {i++}
				if v.y>=centre.y {i+=2} 
				if v.z>=centre.z {i+=4}
				return
			},
		)
	}(){
		cvr <- s
	}
	close(cvr)
	return cvr
}

