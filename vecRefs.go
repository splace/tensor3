package tensor3

type VectorRefs []*Vector

func NewVectorRefs(cs ...BaseType)(vs VectorRefs){
	vs=make(VectorRefs,(len(cs)+2)/3)
	for i:=range(vs){
		v:=NewVector(cs[i*3:]...)
		vs[i]=&v
	}
	return
}

func NewVectorRefsFromIndexes(cs Vectors,indexes ...uint)(vs VectorRefs){
	if len(indexes)==0{
		vs=make(VectorRefs,len(cs))
		for i:=range(cs){
			vs[i]=&cs[i]
		}
	}else{
		vs=make(VectorRefs,len(indexes))
		for i:=range(vs){
			vs[i]=&cs[indexes[i]-1]
		}
	}
	return
}


// rebases a number of vectorrefs to point into a new returned vectors.
func NewVectorsFromVectorRefs(vss ...VectorRefs) Vectors {
	m:=make(map[*Vector]uint)
	for _,vs :=range(vss){
		for _,v:=range(vs) {
			if _,ok:=m[v];!ok{  
				m[v]=uint(len(m))
			}
		}
	}
	nv:=make(Vectors,len(m))
	for vr,index:=range(m){
		nv[index]=*vr
	}
	for _,vs :=range(vss){
		for i,vr:=range(vs) {
			vs[i]=&nv[m[vr]]
		}
	}	
	return nv
}


// TODO find index from pointer, use unsafe? or read text?
func (vsr VectorRefs) Indexes(vs Vectors) (is []uint) {
	is=make([]uint,len(vsr))
	for ir,r:=range(vsr) {
		for i:=range(vs) {
			if &vs[i]==r {is[ir]=uint(i+1);break}
		}
	}
	return is
}


func (vsr VectorRefs) Dereference() (vs Vectors) {
	vs=make(Vectors,len(vsr))
	for i:=range(vs){
		vs[i]=*vsr[i]
	}
	return
}

func (vs Vectors) Dereference(vsr VectorRefs) {
	if len(vs)>len(vsr){
		for i:=range(vsr){
			vs[i]=*vsr[i]
		}
	}else{
		for i:=range(vs){
			vs[i]=*vsr[i]
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
	v.ReduceRefs(vs, (*Vector).Add)
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
	v.ReduceRefs(vs[1:], (*Vector).Max)
	return
}

func (vs VectorRefs) Min() (v Vector) {
	v.Set(*vs[0])
	v.ReduceRefs(vs[1:], (*Vector).Min)
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
		VectorRefsApply(vs, fn, v)
	} else {
		if Hints.ChunkSizeFixed {
			VectorRefsApplyChunked(vs, fn, v, Hints.DefaultChunkSize)
		} else {
			cs := uint(len(vs)) / (Hints.Threads + 1)
			if cs < Hints.DefaultChunkSize {
				cs = Hints.DefaultChunkSize
			}
			VectorRefsApplyChunked(vs, fn, v, cs)
		}
	}
}

func VectorRefsApply(vs VectorRefs, fn func(*Vector, Vector), v Vector) {
	for i := range vs {
		fn(vs[i], v)
	}
}

func VectorRefsApplyChunked(vs VectorRefs, fn func(*Vector, Vector), v Vector, chunkSize uint) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range vectorRefsInChunks(vs, chunkSize) {
		running++
		go func(c VectorRefs) {
			VectorRefsApply(c, fn, v)
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


func (vs Vectors) CrossAllRefs(vs2 VectorRefs) {
	vs.ForAllRefs((*Vector).Cross, vs2)
}


func (vs Vectors) AddAllRefs(vs2 VectorRefs) {
	vs.ForAllRefs((*Vector).Add, vs2)
}


func (vs Vectors) SubtractAllRefs(vs2 VectorRefs) {
	vs.ForAllRefs((*Vector).Subtract, vs2)
}

func (vs Vectors) ProjectAllRefs(vs2 VectorRefs) {
	vs.ForAllRefs((*Vector).Project, vs2)
}

func (vs Vectors) ForAllRefs(fn func(*Vector, Vector), vs2 VectorRefs) {
	if !Parallel {
		vectorsApplyAllRefs(vs, fn, vs2)
	} else {
		if Hints.ChunkSizeFixed {
			vectorsApplyAllRefsChunked(vs, fn, vs2, Hints.DefaultChunkSize)
		} else {
			cs := uint(len(vs)) / (Hints.Threads + 1)
			if cs < Hints.DefaultChunkSize {
				cs = Hints.DefaultChunkSize
			}
			vectorsApplyAllRefsChunked(vs, fn, vs2, cs)
		}
	}
}

func vectorsApplyAllRefs(vs Vectors, fn func(*Vector, Vector), vs2 VectorRefs) {
	for i := range vs {
		fn(&vs[i], *vs2[i])
	}
}

func vectorsApplyAllRefsChunked(vs Vectors, fn func(*Vector, Vector), vs2 VectorRefs, chunkSize uint) {
	done := make(chan struct{}, 1)
	var running uint
	chunks2:=vectorRefsInChunks(vs2, chunkSize)
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


func (vs VectorRefs) CrossAllRefs(vs2 VectorRefs) {
	vs.ForAllRefs((*Vector).Cross, vs2)
}


func (vs VectorRefs) AddAllRefs(vs2 VectorRefs) {
	vs.ForAllRefs((*Vector).Add, vs2)
}


func (vs VectorRefs) SubtractAllRefs(vs2 VectorRefs) {
	vs.ForAllRefs((*Vector).Subtract, vs2)
}

func (vs VectorRefs) ProjectAllRefs(vs2 VectorRefs) {
	vs.ForAllRefs((*Vector).Project, vs2)
}

func (vs VectorRefs) ForAllRefs(fn func(*Vector, Vector), vs2 VectorRefs) {
	if !Parallel {
		vectorRefsApplyAllRefs(vs, fn, vs2)
	} else {
		if Hints.ChunkSizeFixed {
			vectorRefsApplyAllChunkedRefs(vs, fn, vs2, Hints.DefaultChunkSize)
		} else {
			cs := uint(len(vs)) / (Hints.Threads + 1)
			if cs < Hints.DefaultChunkSize {
				cs = Hints.DefaultChunkSize
			}
			vectorRefsApplyAllChunkedRefs(vs, fn, vs2, cs)
		}
	}
}

func vectorRefsApplyAllRefs(vs VectorRefs, fn func(*Vector, Vector), vs2 VectorRefs) {
	for i := range vs {
		fn(vs[i], *vs2[i])
	}
}

func vectorRefsApplyAllChunkedRefs(vs VectorRefs, fn func(*Vector, Vector), vs2 VectorRefs, chunkSize uint) {
	done := make(chan struct{}, 1)
	var running uint
	chunks2:=vectorRefsInChunks(vs2, chunkSize)
	for chunk := range vectorRefsInChunks(vs, chunkSize) {
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


func (vs VectorRefs) CrossAll(vs2 Vectors) {
	vs.ForAll((*Vector).Cross, vs2)
}


func (vs VectorRefs) AddAll(vs2 Vectors) {
	vs.ForAll((*Vector).Add, vs2)
}


func (vs VectorRefs) SubtractAll(vs2 Vectors) {
	vs.ForAll((*Vector).Subtract, vs2)
}

func (vs VectorRefs) ProjectAll(vs2 Vectors) {
	vs.ForAll((*Vector).Project, vs2)
}

func (vs VectorRefs) ForAll(fn func(*Vector, Vector), vs2 Vectors) {
	if !Parallel {
		vectorRefsApplyAll(vs, fn, vs2)
	} else {
		if Hints.ChunkSizeFixed {
			vectorRefsApplyAllChunked(vs, fn, vs2, Hints.DefaultChunkSize)
		} else {
			cs := uint(len(vs)) / (Hints.Threads + 1)
			if cs < Hints.DefaultChunkSize {
				cs = Hints.DefaultChunkSize
			}
			vectorRefsApplyAllChunked(vs, fn, vs2, cs)
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
	chunks2:=vectorsInChunks(vs2, chunkSize)
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


