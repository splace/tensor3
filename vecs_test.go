package tensor3

import "testing"
import "fmt"

func TestVecsPrint(t *testing.T) {
	v := Vectors{*New(1, 2, 3)}
	if fmt.Sprint(v) != "[{1 2 3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsNew(t *testing.T) {
	v := NewVectors(1, 2, 3, 4, 5, 6, 7)
	if fmt.Sprint(v) != "[{1 2 3} {4 5 6} {7 0 0}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCrossLen1(t *testing.T) {
	v := Vectors{*New(1, 2, 3)}
	v.Cross(*New(4, 5, 6))
	if fmt.Sprint(v) != "[{-3 6 -3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCross(t *testing.T) {
	v := Vectors{*New(1, 2, 3), *New(1, 2, 3)}
	v.Cross(*New(4, 5, 6))
	if fmt.Sprint(v) != "[{-3 6 -3} {-3 6 -3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsProduct(t *testing.T) {
	v := Vectors{*New(1, 2, 3), *New(1, 2, 3)}
	v.Product(Identity)
	if fmt.Sprint(v) != "[{1 2 3} {1 2 3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsSum(t *testing.T) {
	vs := Vectors{*New(7, 8, 9), *New(7, 8, 9), *New(7, 8, 9)}
	if fmt.Sprint(vs.Sum()) != "{21 24 27}" {
		t.Error(vs.Sum())
	}
}

func TestVecsMax(t *testing.T) {
	vs := Vectors{*New(4, -1, 11), *New(7, 2, 9), *New(7, 8, 9)}
	if fmt.Sprint(vs.Max()) != "{7 8 11}" {
		t.Error(vs.Max())
	}
}

func TestVecsMaxOne(t *testing.T) {
	vs := Vectors{*New(1, -1, 1)}
	if fmt.Sprint(vs.Max()) != "{1 -1 1}" {
		t.Error(vs.Max())
	}
}

func TestVecsMaxNone(t *testing.T) {
	vs := Vectors{}
	defer func() {
		r := recover()
		if r == nil{
			t.Error("Expected error not present.")
		}
	}()
	_ = vs.Max()
}

func TestVecsMin(t *testing.T) {
	vs := Vectors{*New(4, -1, 11), *New(7, 2, 9), *New(7, 8, 9)}
	if fmt.Sprint(vs.Min()) != "{4 -1 9}" {
		t.Error(vs.Min())
	}
}

func TestVecsMiddle(t *testing.T) {
	vs := Vectors{*New(4, -1, 11), *New(7, 2, 9), *New(7, 8, 9)}
	if fmt.Sprint(vs.Middle()) != "{5.5 3.5 10}" {
		t.Error(vs.Middle())
	}
}

func TestVecsInterpolate(t *testing.T) {
	vs := Vectors{*New(7, 8, 9), *New(7, 8, 9), *New(7, 8, 9)}
	vs.Interpolate(Vector{-2, 1, -1}, 0.5)
	if fmt.Sprint(vs) != "[{2.5 4.5 4} {2.5 4.5 4} {2.5 4.5 4}]" {
		t.Error(vs)
	}
}

func TestVecsExtrapolate(t *testing.T) {
	vs := Vectors{*New(7, 8, 9), *New(7, 8, 9), *New(7, 8, 9)}
	vs.Interpolate(*New(-2, 1, -1), 2)
	if fmt.Sprint(vs) != "[{16 15 19} {16 15 19} {16 15 19}]" {
		t.Error(vs)
	}
}

func TestVecsProductT(t *testing.T) {
	v := Vectors{*New(1, 2, 3), *New(1, 2, 3)}
	v.ProductT(Identity)
	if fmt.Sprint(v) != "[{1 2 3} {1 2 3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCrossChunked1(t *testing.T) {
	Parallel = true
	defer func(d int) {
		Parallel = false
		Hints.DefaultChunkSize = d
	}(Hints.DefaultChunkSize)
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 1
	v := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9), *New(10, 11, 12)}
	//	v=append(v,Vector{1, 2, 3})
	v.Add(*New(1, 1, 1))
	if fmt.Sprint(v) != "[{2 3 4} {5 6 7} {8 9 10} {11 12 13}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCrossChunked2(t *testing.T) {
	Parallel = true
	defer func(d int) {
		Parallel = false
		Hints.DefaultChunkSize = d
	}(Hints.DefaultChunkSize)
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 2
	v := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9), *New(10, 11, 12)}
	//	v=append(v,Vector{1, 2, 3})
	v.Add(*New(1, 1, 1))
	if fmt.Sprint(v) != "[{2 3 4} {5 6 7} {8 9 10} {11 12 13}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCrossChunked3(t *testing.T) {
	Parallel = true
	defer func(d int) {
		Parallel = false
		Hints.DefaultChunkSize = d
	}(Hints.DefaultChunkSize)
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 2
	v := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9), *New(10, 11, 12), *New(13, 14, 15), *New(16, 17, 18)}
	//	v=append(v,Vector{1, 2, 3})
	v.Add(*New(1, 1, 1))
	if fmt.Sprint(v) != "[{2 3 4} {5 6 7} {8 9 10} {11 12 13} {14 15 16} {17 18 19}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsAddVecs(t *testing.T) {
	vs := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9)}
	vs2 := Vectors{*New(9, 8, 7), *New(6, 5, 4), *New(3, 2, 1)}
	vs.AddAll(vs2)
	if fmt.Sprint(vs[0], vs[1], vs[2]) != "{10 10 10} {10 10 10} {10 10 10}" {
		t.Error(fmt.Sprint(vs[0], vs[1], vs[2]))
	}
}

func TestVecsCrossVecs(t *testing.T) {
	vs := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9)}
	vs2 := Vectors{*New(9, 8, 7), *New(6, 5, 4), *New(3, 2, 1)}
	vs.CrossAll(vs2)
	if fmt.Sprint(vs[0], vs[1], vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}" {
		t.Error(fmt.Sprint(vs[0], vs[1], vs[2]))
	}
}


func TestVecsSlicesInChunks(t *testing.T) {
	Hints.ChunkSizeFixed = true
	defer func(dcs int) {
		Hints.ChunkSizeFixed = false
		Hints.DefaultChunkSize = dcs 
	}(Hints.DefaultChunkSize)
	Hints.DefaultChunkSize = 2

	vs := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9), *New(10, 11, 12), *New(13, 14, 15)}
	var vs2 [][]Vectors
	
	for vss:=range vectorSlicesInChunks(vs,10,1,1,true){
		vs2=append(vs2,vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3}] [{4 5 6}] [{7 8 9}] [{10 11 12}] [{13 14 15}]]]" {
		t.Error(fmt.Println(vs2))
	}
	vs2=vs2[:0]
	for vss:=range vectorSlicesInChunks(vs,10,2,1,false){
		vs2=append(vs2,vss)
	}

	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6}] [{4 5 6} {7 8 9}] [{7 8 9} {10 11 12}] [{10 11 12} {13 14 15}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2=vs2[:0]
	for vss:=range vectorSlicesInChunks(vs,10,3,1,true){
		vs2=append(vs2,vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6} {7 8 9}] [{4 5 6} {7 8 9} {10 11 12}] [{7 8 9} {10 11 12} {13 14 15}] [{10 11 12} {13 14 15} {1 2 3}] [{13 14 15} {1 2 3} {4 5 6}]]]" {
		t.Error(fmt.Println(vs2))
	}


	vs2=vs2[:0]
	for vss:=range vectorSlicesInChunks(vs,1,1,1,true){
		vs2=append(vs2,vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3}]] [[{4 5 6}]] [[{7 8 9}]] [[{10 11 12}]] [[{13 14 15}]]]" {
		t.Error(fmt.Println(vs2))
	}


	vs2=vs2[:0]
	for vss:=range vectorSlicesInChunks(vs,2,2,1,true){
		vs2=append(vs2,vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6}] [{4 5 6} {7 8 9}]] [[{7 8 9} {10 11 12}] [{10 11 12} {13 14 15}] [{13 14 15} {1 2 3}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2=vs2[:0]
	for vss:=range vectorSlicesInChunks(vs,4,2,1,false){
		vs2=append(vs2,vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6}] [{4 5 6} {7 8 9}] [{7 8 9} {10 11 12}] [{10 11 12} {13 14 15}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2=vs2[:0]
	for vss:=range vectorSlicesInChunks(vs,3,1,1,false){
		vs2=append(vs2,vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3}] [{4 5 6}] [{7 8 9}]] [[{10 11 12}] [{13 14 15}]]]" {
		t.Error(fmt.Println(vs2))
	}

}

func TestVecsSlicesStridingInChunks(t *testing.T) {
	Hints.ChunkSizeFixed = true
	defer func(dcs int) {
		Hints.ChunkSizeFixed = false
		Hints.DefaultChunkSize = dcs 
	}(Hints.DefaultChunkSize)
	Hints.DefaultChunkSize = 2

	vs := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9), *New(10, 11, 12), *New(13, 14, 15)}
	var vs2 [][]Vectors
	
	for vss:=range vectorSlicesInChunks(vs,10,1,2,false){
		vs2=append(vs2,vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3}]] [[{7 8 9}]] [[{13 14 15}]]]" {
		t.Error(fmt.Println(vs2))
	}

}



func BenchmarkVecsSum(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 100000)
	for i := range vs {
		vs[i] = *New(1, 2, 3)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Sum()
	}
}

func BenchmarkVecsSumParallel(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 100000)
	for i := range vs {
		vs[i] = *New(1, 2, 3)
	}
	Parallel = true
	Hints.ChunkSizeFixed = true
	defer func() {
		Parallel = false
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Sum()
	}
}

func BenchmarkVecsCross(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 100000)
	for i := range vs {
		vs[i] = *New(1, 2, 3)
	}
	v := Vector{9, 8, 7}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Cross(v)
	}

}

func BenchmarkVecsCrossParallel(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 100000)
	for i := range vs {
		vs[i] = *New(1, 2, 3)
	}
	v := *New(9, 8, 7)
	Parallel = true
	Hints.ChunkSizeFixed = true
	defer func() {
		Parallel = false
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Cross(v)
	}

}

func BenchmarkVecsProduct(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 100000)
	for i := range vs {
		vs[i] = *New(1, 2, 3)
	}
	m := Matrix{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Product(m)
	}

}

func BenchmarkVecsProductParallel(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 100000)
	for i := range vs {
		vs[i] = *New(1, 2, 3)
	}
	m := Matrix{}
	Parallel = true
	Hints.ChunkSizeFixed = true
	defer func() {
		Parallel = false
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Product(m)
	}

}

/*  Hal3 Thu 24 May 21:42:01 BST 2018 go version go1.6.2 linux/amd64
=== RUN   TestMatrixProductT
--- PASS: TestMatrixProductT (0.00s)
=== RUN   TestMatrixTProduct
--- PASS: TestMatrixTProduct (0.00s)
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
=== RUN   TestMatrixProductIndentity
--- PASS: TestMatrixProductIndentity (0.00s)
=== RUN   TestMatrixMultiply
--- PASS: TestMatrixMultiply (0.00s)
=== RUN   TestMatrixAggregate
--- PASS: TestMatrixAggregate (0.00s)
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
=== RUN   TestVecRefsSelect
--- PASS: TestVecRefsSelect (0.00s)
=== RUN   TestVecRefsSplit
--- PASS: TestVecRefsSplit (0.00s)
=== RUN   TestVecPrint
--- PASS: TestVecPrint (0.00s)
=== RUN   TestNew
--- PASS: TestNew (0.00s)
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
--- PASS: TestVecsMaxNone (0.00s)
=== RUN   TestVecsMin
--- PASS: TestVecsMin (0.00s)
=== RUN   TestVecsInterpolate
--- PASS: TestVecsInterpolate (0.00s)
=== RUN   TestVecsExtrapolate
--- PASS: TestVecsExtrapolate (0.00s)
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
=== RUN   TestVecsSlicesInChunks
--- PASS: TestVecsSlicesInChunks (0.00s)
=== RUN   TestVecsSlicesStridingInChunks
[[[{1 2 3}] []]]
--- FAIL: TestVecsSlicesStridingInChunks (0.00s)
	vecs_test.go:259: 17 <nil>
=== RUN   ExampleTextEncodingVector
--- PASS: ExampleTextEncodingVector (0.00s)
=== RUN   ExampleTextEncodingMatrix
--- PASS: ExampleTextEncodingMatrix (0.00s)
=== RUN   ExampleBounds
--- PASS: ExampleBounds (0.00s)
=== RUN   ExampleUnitify
--- PASS: ExampleUnitify (0.00s)
=== RUN   ExampleForEachVector
--- PASS: ExampleForEachVector (0.00s)
FAIL
exit status 1
FAIL	_/home/simon/Dropbox/github/working/tensor3	0.006s
Thu 24 May 21:42:03 BST 2018
*/

