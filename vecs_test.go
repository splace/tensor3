package tensor3

import "testing"
import "fmt"

func TestVecsPrint(t *testing.T) {
	v := Vectors{Vector{1, 2, 3}}
	if fmt.Sprint(v) != "[{1 2 3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsNew(t *testing.T) {
	v := NewVectors(1, 2, 3, 4, 5, 6,7)
	if fmt.Sprint(v) != "[{1 2 3} {4 5 6} {7 0 0}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCrossLen1(t *testing.T) {
	v := Vectors{Vector{1, 2, 3}}
	v.Cross(Vector{4, 5, 6})
	if fmt.Sprint(v) != "[{-3 6 -3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCross(t *testing.T) {
	v := Vectors{Vector{1, 2, 3}, Vector{1, 2, 3}}
	v.Cross(Vector{4, 5, 6})
	if fmt.Sprint(v) != "[{-3 6 -3} {-3 6 -3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsProduct(t *testing.T) {
	v := Vectors{Vector{1, 2, 3}, Vector{1, 2, 3}}
	v.Product(Identity)
	if fmt.Sprint(v) != "[{1 2 3} {1 2 3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsSum(t *testing.T) {
	vs := Vectors{Vector{7, 8, 9}, Vector{7, 8, 9}, Vector{7, 8, 9}}
	if fmt.Sprint(vs.Sum()) != "{21 24 27}" {
		t.Error(vs.Sum())
	}
}

func TestVecsMax(t *testing.T) {
	vs := Vectors{Vector{4, -1, 11}, Vector{7, 2, 9}, Vector{7, 8, 9}}
	if fmt.Sprint(vs.Max()) != "{7 8 11}" {
		t.Error(vs.Max())
	}
}

func TestVecsMaxOne(t *testing.T) {
	vs := Vectors{Vector{1, -1, 1}}
	if fmt.Sprint(vs.Max()) != "{1 -1 1}" {
		t.Error(vs.Max())
	}
}

func TestVecsMaxNone(t *testing.T) {
	vs := Vectors{}
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("As expected:", r)
		} else {
			t.Error("Error expected.")
		}
	}()
	_ = vs.Max()
}

func TestVecsMin(t *testing.T) {
	vs := Vectors{Vector{4, -1, 11}, Vector{7, 2, 9}, Vector{7, 8, 9}}
	if fmt.Sprint(vs.Min()) != "{4 -1 9}" {
		t.Error(vs.Min())
	}
}

func TestVecsInterpolate(t *testing.T) {
	vs := Vectors{Vector{7, 8, 9}, Vector{7, 8, 9}, Vector{7, 8, 9}}
	vs.Interpolate(Vector{-2, 1, -1}, 0.5)
	if fmt.Sprint(vs) != "[{2.5 4.5 4} {2.5 4.5 4} {2.5 4.5 4}]" {
		t.Error(vs)
	}
}

func TestVecsProductT(t *testing.T) {
	v := Vectors{Vector{1, 2, 3}, Vector{1, 2, 3}}
	v.ProductT(Identity)
	if fmt.Sprint(v) != "[{1 2 3} {1 2 3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCrossChunked1(t *testing.T) {
	Parallel = true
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 1
	v := Vectors{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}, Vector{10, 11, 12}}
	//	v=append(v,Vector{1, 2, 3})
	v.Add(Vector{1, 1, 1})
	if fmt.Sprint(v) != "[{2 3 4} {5 6 7} {8 9 10} {11 12 13}]" {
		t.Error(fmt.Sprint(v))
	}
	Parallel = false
}

func TestVecsCrossChunked2(t *testing.T) {
	Parallel = true
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 2
	v := Vectors{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}, Vector{10, 11, 12}}
	//	v=append(v,Vector{1, 2, 3})
	v.Add(Vector{1, 1, 1})
	if fmt.Sprint(v) != "[{2 3 4} {5 6 7} {8 9 10} {11 12 13}]" {
		t.Error(fmt.Sprint(v))
	}
	Parallel = false
}

func TestVecsCrossChunked3(t *testing.T) {
	Parallel = true
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 2
	v := Vectors{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}, Vector{10, 11, 12}, Vector{13, 14, 15}, Vector{16, 17, 18}}
	//	v=append(v,Vector{1, 2, 3})
	v.Add(Vector{1, 1, 1})
	if fmt.Sprint(v) != "[{2 3 4} {5 6 7} {8 9 10} {11 12 13} {14 15 16} {17 18 19}]" {
		t.Error(fmt.Sprint(v))
	}
	Parallel = false
}

func BenchmarkVecsCross(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 1000000)
	for i := range vs {
		vs[i] = Vector{1, 2, 3}
	}
	v := Vector{9, 8, 7}
	Parallel = false
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Cross(v)
	}

}

func BenchmarkVecsCrossParallel(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 1000000)
	for i := range vs {
		vs[i] = Vector{1, 2, 3}
	}
	v := Vector{9, 8, 7}
	Parallel = true
	Hints.ChunkSizeFixed = true
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Cross(v)
	}

}

func BenchmarkVecsProduct(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 1000000)
	for i := range vs {
		vs[i] = Vector{1, 2, 3}
	}
	m := Matrix{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Product(m)
	}

}

func BenchmarkVecsProductParallel(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 1000000)
	for i := range vs {
		vs[i] = Vector{1, 2, 3}
	}
	m := Matrix{}
	Parallel = true
	Hints.ChunkSizeFixed = true
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Product(m)
	}

}

/*  Hal3 Mon 6 Mar 23:51:24 GMT 2017 go version go1.6.2 linux/amd64
PASS
BenchmarkVecsCross-2        	     100	  14147401 ns/op
BenchmarkVecsCrossParallel-2	     100	  13179012 ns/op
ok  	_/home/simon/Dropbox/github/working/tensor3	2.905s
Mon 6 Mar 23:51:28 GMT 2017
*/

/*  Hal3 Tue 7 Mar 00:09:28 GMT 2017 go version go1.6.2 linux/amd64
PASS
BenchmarkVecsCross-2        	     100	  13722238 ns/op
BenchmarkVecsCrossParallel-2	     100	  13090129 ns/op
ok  	_/home/simon/Dropbox/github/working/tensor3	2.837s
Tue 7 Mar 00:09:32 GMT 2017
*/

/*  Hal3 Tue 7 Mar 00:39:43 GMT 2017 go version go1.6.2 linux/amd64
PASS
BenchmarkVecsCross-2          	     100	  14792161 ns/op
BenchmarkVecsCrossParallel-2  	     100	  14016910 ns/op
BenchmarkVecsProductParallel-2	     100	  14267832 ns/op
ok  	_/home/simon/Dropbox/github/working/tensor3	4.561s
Tue 7 Mar 00:39:49 GMT 2017
*/
/*  Hal3 Tue 7 Mar 00:42:46 GMT 2017 go version go1.6.2 linux/amd64
PASS
BenchmarkVecsCross-2          	     100	  14772536 ns/op
BenchmarkVecsCrossParallel-2  	     100	  13781806 ns/op
BenchmarkVecsProduct-2        	     100	  14155553 ns/op
BenchmarkVecsProductParallel-2	     100	  14211055 ns/op
ok  	_/home/simon/Dropbox/github/working/tensor3	6.073s
Tue 7 Mar 00:42:53 GMT 2017
*/
/*  Hal3 Tue 7 Mar 00:52:00 GMT 2017  go version go1.8 linux/amd64

BenchmarkVecsCross-2             	     100	  13334312 ns/op
BenchmarkVecsCrossParallel-2     	     100	  12970468 ns/op
BenchmarkVecsProduct-2           	     100	  13046767 ns/op
BenchmarkVecsProductParallel-2   	     100	  13137967 ns/op
PASS
ok  	_/home/simon/Dropbox/github/working/tensor3	5.553s
Tue 7 Mar 00:52:07 GMT 2017
*/
