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

	vs := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9)}
	
	// loop to receive Vectors but for these settings only one actual chunk produced.
	for vss:=range vectorSlicesInChunks(vs,chunkSize(len(vs)),1,1,true){
		if fmt.Sprint(vss) != "[[{1 2 3}] [{4 5 6}] [{7 8 9}]]" {
			t.Error(fmt.Println(vss))
		}
	}
	for vss:=range vectorSlicesInChunks(vs,chunkSize(len(vs)),2,1,false){
		if fmt.Sprint(vss) != "[[{1 2 3} {4 5 6}] [{4 5 6} {7 8 9}]]" {
			t.Error(fmt.Println(vss))
		}
	}
	for vss:=range vectorSlicesInChunks(vs,chunkSize(len(vs)),3,1,true){
		if fmt.Sprint(vss) != "[[{1 2 3} {4 5 6} {7 8 9}] [{4 5 6} {7 8 9} {1 2 3}] [{7 8 9} {1 2 3} {4 5 6}]]" {
			t.Error(fmt.Println(vss))
		}
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

/* benchmark: "Vecs" hal3 Thu 24 May 16:09:48 BST 2018 go version go1.10.2 linux/amd64
goos: linux
goarch: amd64
BenchmarkVecsSum-2               	    2000	    965165 ns/op
BenchmarkVecsSumParallel-2       	    2000	    900285 ns/op
BenchmarkVecsCross-2             	    1000	   1417655 ns/op
BenchmarkVecsCrossParallel-2     	    1000	   1429321 ns/op
BenchmarkVecsProduct-2           	    1000	   1951688 ns/op
BenchmarkVecsProductParallel-2   	    1000	   1940458 ns/op
PASS
ok  	_/run/media/simon/6a5530c2-1442-4e9b-b35f-3db0c9a6984c/home/simon/Dropbox/github/working/tensor3	11.463s
Thu 24 May 16:10:00 BST 2018
*/

