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


// rebases, maintaining values, a number of vectorrefs to point into a new returned vectors.
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
	vs.ForAllRefs(vrs,(*Vector).Cross)
}


func (vs Vectors) AddAllRefs(vrs VectorRefs) {
	vs.ForAllRefs(vrs,(*Vector).Add)
}


func (vs Vectors) SubtractAllRefs(vrs VectorRefs) {
	vs.ForAllRefs(vrs,(*Vector).Subtract)
}

func (vs Vectors) ProjectAllRefs(vrs VectorRefs) {
	vs.ForAllRefs(vrs,(*Vector).Project)
}

func (vs Vectors) ForAllRefs(vrs VectorRefs,fn func(*Vector, Vector)) {
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
	chunks2:=vectorRefsInChunks(vrs, chunkSize)
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
	vrs.ForAllRefs(vrs2,(*Vector).Cross)
}


func (vrs VectorRefs) AddAllRefs(vrs2 VectorRefs) {
	vrs.ForAllRefs(vrs2,(*Vector).Add)
}


func (vrs VectorRefs) SubtractAllRefs(vrs2 VectorRefs) {
	vrs.ForAllRefs(vrs2,(*Vector).Subtract)
}

func (vrs VectorRefs) ProjectAllRefs(vrs2 VectorRefs) {
	vrs.ForAllRefs(vrs2,(*Vector).Project)
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
	chunks2:=vectorRefsInChunks(vrs2, chunkSize)
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
	vrs.ForAll(vs,(*Vector).Cross)
}


func (vrs VectorRefs) AddAll(vs Vectors) {
	vrs.ForAll(vs,(*Vector).Add)
}


func (vrs VectorRefs) SubtractAll(vs Vectors) {
	vrs.ForAll(vs,(*Vector).Subtract)
}

func (vrs VectorRefs) ProjectAll(vs Vectors) {
	vrs.ForAll(vs,(*Vector).Project)
}

func (vrs VectorRefs) ForAll(vs Vectors,fn func(*Vector, Vector)) {
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


/*  Hal3 Sat 15 Apr 00:12:56 BST 2017 go version go1.6.2 linux/amd64
=== RUN   TestMatrixProductT
--- PASS: TestMatrixProductT (0.00s)
=== RUN   TestMatrixTProductT
--- PASS: TestMatrixTProductT (0.00s)
=== RUN   TestMatrixIdentityPrint
--- PASS: TestMatrixIdentityPrint (0.00s)
=== RUN   TestMatrixNew
--- PASS: TestMatrixNew (0.00s)
=== RUN   TestMatrixPrint
--- PASS: TestMatrixPrint (0.00s)
=== RUN   TestMatrixAdd
--- PASS: TestMatrixAdd (0.00s)
=== RUN   TestMatrixDeterminant
--- PASS: TestMatrixDeterminant (0.00s)
=== RUN   TestMatrixInvert
--- PASS: TestMatrixInvert (0.00s)
=== RUN   TestMatrixProduct
--- PASS: TestMatrixProduct (0.00s)
=== RUN   TestMatrixProductDet
--- PASS: TestMatrixProductDet (0.00s)
=== RUN   TestMatrixInvertDet
--- PASS: TestMatrixInvertDet (0.00s)
=== RUN   TestMatrixTProduct
--- PASS: TestMatrixTProduct (0.00s)
=== RUN   TestMatrixProductIndentity
--- PASS: TestMatrixProductIndentity (0.00s)
=== RUN   TestMatrixMultiply
--- PASS: TestMatrixMultiply (0.00s)
=== RUN   TestMatrixReduce
--- PASS: TestMatrixReduce (0.00s)
=== RUN   TestMatrixApplyXAdd
--- PASS: TestMatrixApplyXAdd (0.00s)
=== RUN   TestMatrixApplyZAdd
--- PASS: TestMatrixApplyZAdd (0.00s)
=== RUN   TestMatrixApplyComponentWiseAxesAdd
--- PASS: TestMatrixApplyComponentWiseAxesAdd (0.00s)
=== RUN   TestMatrixApplyComponentWiseAxesCross
--- PASS: TestMatrixApplyComponentWiseAxesCross (0.00s)
=== RUN   TestMatsPrint
--- PASS: TestMatsPrint (0.00s)
=== RUN   TestMatsProduct1
--- PASS: TestMatsProduct1 (0.00s)
=== RUN   TestMatsSum
--- PASS: TestMatsSum (0.00s)
=== RUN   TestMatsApplyComponents
--- PASS: TestMatsApplyComponents (0.00s)
=== RUN   TestMatsApplyComponentY
--- PASS: TestMatsApplyComponentY (0.00s)
=== RUN   TestMatsApplyComponentsXY
--- PASS: TestMatsApplyComponentsXY (0.00s)
=== RUN   TestVecRefsNewFromIndexes
--- PASS: TestVecRefsNewFromIndexes (0.00s)
=== RUN   TestVecRefsNewFromEmplyIndexes
--- PASS: TestVecRefsNewFromEmplyIndexes (0.00s)
=== RUN   TestVecRefsDereferenceNew
--- PASS: TestVecRefsDereferenceNew (0.00s)
=== RUN   TestVecRefsDereference
--- PASS: TestVecRefsDereference (0.00s)
=== RUN   TestVecsFromVectorRefs
--- PASS: TestVecsFromVectorRefs (0.00s)
=== RUN   TestVecsIndexesFromVectorRefs
--- PASS: TestVecsIndexesFromVectorRefs (0.00s)
=== RUN   TestVecRefsPrint
--- PASS: TestVecRefsPrint (0.00s)
=== RUN   TestVecRefsSum
--- PASS: TestVecRefsSum (0.00s)
=== RUN   TestVecRefsAddRefs
--- PASS: TestVecRefsAddRefs (0.00s)
=== RUN   TestVecRefsCrossRefs
--- PASS: TestVecRefsCrossRefs (0.00s)
=== RUN   TestVecsAddVecRefs
--- PASS: TestVecsAddVecRefs (0.00s)
=== RUN   TestVecsCrossVecRefs
--- PASS: TestVecsCrossVecRefs (0.00s)
=== RUN   TestVecsRefsAddVecs
--- PASS: TestVecsRefsAddVecs (0.00s)
=== RUN   TestVecRefsCrossVecs
--- PASS: TestVecRefsCrossVecs (0.00s)
=== RUN   TestVecsRefsAddVecRefs
--- PASS: TestVecsRefsAddVecRefs (0.00s)
=== RUN   TestVecRefsCrossVecRefs
--- PASS: TestVecRefsCrossVecRefs (0.00s)
=== RUN   TestVecPrint
--- PASS: TestVecPrint (0.00s)
=== RUN   TestNewVector
--- PASS: TestNewVector (0.00s)
=== RUN   TestNewVectorPrint
--- PASS: TestNewVectorPrint (0.00s)
=== RUN   TestVecDot
--- PASS: TestVecDot (0.00s)
=== RUN   TestVecAdd
--- PASS: TestVecAdd (0.00s)
=== RUN   TestVecSubtract
--- PASS: TestVecSubtract (0.00s)
=== RUN   TestVecCross
--- PASS: TestVecCross (0.00s)
=== RUN   TestVecLengthLength
--- PASS: TestVecLengthLength (0.00s)
=== RUN   TestVecProduct
--- PASS: TestVecProduct (0.00s)
=== RUN   TestVecProductT
--- PASS: TestVecProductT (0.00s)
=== RUN   TestVecProject
--- PASS: TestVecProject (0.00s)
=== RUN   TestVecLongestAxis
--- PASS: TestVecLongestAxis (0.00s)
=== RUN   TestVecShortestAxis
--- PASS: TestVecShortestAxis (0.00s)
=== RUN   TestVecMax
--- PASS: TestVecMax (0.00s)
=== RUN   TestVecMin
--- PASS: TestVecMin (0.00s)
=== RUN   TestVecMid
--- PASS: TestVecMid (0.00s)
=== RUN   TestVecInterpolate
--- PASS: TestVecInterpolate (0.00s)
=== RUN   TestVecApplyRunning
--- PASS: TestVecApplyRunning (0.00s)
=== RUN   TestVecApplyForAll
--- PASS: TestVecApplyForAll (0.00s)
=== RUN   TestVecsPrint
--- PASS: TestVecsPrint (0.00s)
=== RUN   TestVecsNew
--- PASS: TestVecsNew (0.00s)
=== RUN   TestVecsCrossLen1
--- PASS: TestVecsCrossLen1 (0.00s)
=== RUN   TestVecsCross
--- PASS: TestVecsCross (0.00s)
=== RUN   TestVecsProduct
--- PASS: TestVecsProduct (0.00s)
=== RUN   TestVecsSum
--- PASS: TestVecsSum (0.00s)
=== RUN   TestVecsMax
--- PASS: TestVecsMax (0.00s)
=== RUN   TestVecsMaxOne
--- PASS: TestVecsMaxOne (0.00s)
=== RUN   TestVecsMaxNone
As expected: runtime error: index out of range
--- PASS: TestVecsMaxNone (0.00s)
=== RUN   TestVecsMin
--- PASS: TestVecsMin (0.00s)
=== RUN   TestVecsInterpolate
--- PASS: TestVecsInterpolate (0.00s)
=== RUN   TestVecsProductT
--- PASS: TestVecsProductT (0.00s)
=== RUN   TestVecsCrossChunked1
--- PASS: TestVecsCrossChunked1 (0.00s)
=== RUN   TestVecsCrossChunked2
--- PASS: TestVecsCrossChunked2 (0.00s)
=== RUN   TestVecsCrossChunked3
--- PASS: TestVecsCrossChunked3 (0.00s)
=== RUN   TestVecsAddVecs
--- PASS: TestVecsAddVecs (0.00s)
=== RUN   TestVecsCrossVecs
--- PASS: TestVecsCrossVecs (0.00s)
PASS
ok  	_/home/simon/Dropbox/github/working/tensor3	0.006s
Sat 15 Apr 00:12:58 BST 2017
*/

