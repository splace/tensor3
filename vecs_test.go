package tensor3

import "testing"
import "fmt"

func TestVecsPrint(t *testing.T) {
	v := Vectors{NewVector(1, 2, 3)}
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
	v := Vectors{NewVector(1, 2, 3)}
	v.Cross(NewVector(4, 5, 6))
	if fmt.Sprint(v) != "[{-3 6 -3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCross(t *testing.T) {
	v := Vectors{NewVector(1, 2, 3), NewVector(1, 2, 3)}
	v.Cross(NewVector(4, 5, 6))
	if fmt.Sprint(v) != "[{-3 6 -3} {-3 6 -3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsProduct(t *testing.T) {
	v := Vectors{NewVector(1, 2, 3), NewVector(1, 2, 3)}
	v.Product(Identity)
	if fmt.Sprint(v) != "[{1 2 3} {1 2 3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsSum(t *testing.T) {
	vs := Vectors{NewVector(7, 8, 9), NewVector(7, 8, 9), NewVector(7, 8, 9)}
	if fmt.Sprint(vs.Sum()) != "{21 24 27}" {
		t.Error(vs.Sum())
	}
}

func TestVecsMax(t *testing.T) {
	vs := Vectors{NewVector(4, -1, 11), NewVector(7, 2, 9), NewVector(7, 8, 9)}
	if fmt.Sprint(vs.Max()) != "{7 8 11}" {
		t.Error(vs.Max())
	}
}

func TestVecsMaxOne(t *testing.T) {
	vs := Vectors{NewVector(1, -1, 1)}
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
	vs := Vectors{NewVector(4, -1, 11), NewVector(7, 2, 9), NewVector(7, 8, 9)}
	if fmt.Sprint(vs.Min()) != "{4 -1 9}" {
		t.Error(vs.Min())
	}
}

func TestVecsInterpolate(t *testing.T) {
	vs := Vectors{NewVector(7, 8, 9), NewVector(7, 8, 9), NewVector(7, 8, 9)}
//	vs.Interpolate(Vector{-2, 1, -1}, 0.5)
//	if fmt.Sprint(vs) != "[{2.5 4.5 4} {2.5 4.5 4} {2.5 4.5 4}]" {
	vs.Interpolate(NewVector(-2, 1, -1), 2)
	if fmt.Sprint(vs) != "[{16 15 19} {16 15 19} {16 15 19}]" {
		t.Error(vs)
	}
}

func TestVecsProductT(t *testing.T) {
	v := Vectors{NewVector(1, 2, 3), NewVector(1, 2, 3)}
	v.ProductT(Identity)
	if fmt.Sprint(v) != "[{1 2 3} {1 2 3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCrossChunked1(t *testing.T) {
	Parallel = true
	defer func() {
		Parallel=false
	}()
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 1
	v := Vectors{NewVector(1, 2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9), NewVector(10, 11, 12)}
	//	v=append(v,Vector{1, 2, 3})
	v.Add(NewVector(1, 1, 1))
	if fmt.Sprint(v) != "[{2 3 4} {5 6 7} {8 9 10} {11 12 13}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCrossChunked2(t *testing.T) {
	Parallel = true
	defer func() {
		Parallel=false
	}()
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 2
	v := Vectors{NewVector(1, 2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9), NewVector(10, 11, 12)}
	//	v=append(v,Vector{1, 2, 3})
	v.Add(NewVector(1, 1, 1))
	if fmt.Sprint(v) != "[{2 3 4} {5 6 7} {8 9 10} {11 12 13}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCrossChunked3(t *testing.T) {
	Parallel = true
	defer func() {
		Parallel=false
	}()
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 2
	v := Vectors{NewVector(1, 2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9), NewVector(10, 11, 12), NewVector(13, 14, 15),NewVector(16, 17, 18)}
	//	v=append(v,Vector{1, 2, 3})
	v.Add(NewVector(1, 1, 1))
	if fmt.Sprint(v) != "[{2 3 4} {5 6 7} {8 9 10} {11 12 13} {14 15 16} {17 18 19}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsAddVecs(t *testing.T) {
	vs := Vectors{NewVector(1,2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9)}
	vs2 := Vectors{NewVector(9, 8, 7), NewVector(6, 5, 4), NewVector(3, 2, 1)}
	vs.AddAll(vs2)
	if fmt.Sprint(vs[0],vs[1],vs[2]) != "{10 10 10} {10 10 10} {10 10 10}"{
		t.Error(fmt.Sprint(vs[0],vs[1],vs[2]))
	}
}

func TestVecsCrossVecs(t *testing.T) {
	vs := Vectors{NewVector(1,2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9)}
	vs2 := Vectors{NewVector(9, 8, 7), NewVector(6, 5, 4), NewVector(3, 2, 1)}
	vs.CrossAll(vs2)
	if fmt.Sprint(vs[0],vs[1],vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}"{
		t.Error(fmt.Sprint(vs[0],vs[1],vs[2]))
	}
}



func BenchmarkVecsSum(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 1000000)
	for i := range vs {
		vs[i] = NewVector(1, 2, 3)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Sum()
	}
}


func BenchmarkVecsSumParallel(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 1000000)
	for i := range vs {
		vs[i] = NewVector(1, 2, 3)
	}
	Parallel = true
	defer func() {
		Parallel=false
	}()
	Hints.ChunkSizeFixed = true
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Sum()
	}
}

func BenchmarkVecsCross(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 1000000)
	for i := range vs {
		vs[i] = NewVector(1, 2, 3)
	}
	v := Vector{9, 8, 7}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Cross(v)
	}

}

func BenchmarkVecsCrossParallel(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 1000000)
	for i := range vs {
		vs[i] = NewVector(1, 2, 3)
	}
	v := NewVector(9, 8, 7)
	Parallel = true
	defer func() {
		Parallel=false
	}()
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
		vs[i] = NewVector(1, 2, 3)
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
		vs[i] = NewVector(1, 2, 3)
	}
	m := Matrix{}
	Parallel = true
	defer func() {
		Parallel=false
	}()
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
*//*  Hal3 Sat 15 Apr 00:28:58 BST 2017 go version go1.6.2 linux/amd64
PASS
BenchmarkVecsSum-2            	     100	  10705533 ns/op
BenchmarkVecsSumParallel-2    	     100	  10577862 ns/op
BenchmarkVecsCross-2          	     100	  13915981 ns/op
BenchmarkVecsCrossParallel-2  	     100	  13618217 ns/op
BenchmarkVecsProduct-2        	     100	  21823835 ns/op
BenchmarkVecsProductParallel-2	     100	  13131134 ns/op
ok  	_/home/simon/Dropbox/github/working/tensor3	8.850s
Sat 15 Apr 00:29:08 BST 2017
*/
/*  Hal3 Sat 15 Apr 00:31:18 BST 2017  go version go1.8 linux/amd64

BenchmarkVecsSum-2               	     200	   8384060 ns/op
BenchmarkVecsSumParallel-2       	     200	   8307832 ns/op
BenchmarkVecsCross-2             	     100	  14350239 ns/op
BenchmarkVecsCrossParallel-2     	     100	  13895992 ns/op
BenchmarkVecsProduct-2           	     100	  19148612 ns/op
BenchmarkVecsProductParallel-2   	     100	  14231442 ns/op
PASS
ok  	_/home/simon/Dropbox/github/working/tensor3	11.329s
Sat 15 Apr 00:31:30 BST 2017
*/
/*  Hal3 Tue 30 May 14:20:08 BST 2017 go version go1.6.2 linux/amd64
PASS
BenchmarkVecsSum-2            	     200	   9785954 ns/op
BenchmarkVecsSumParallel-2    	     200	   9807838 ns/op
BenchmarkVecsCross-2          	     100	  14532405 ns/op
BenchmarkVecsCrossParallel-2  	     100	  14020157 ns/op
BenchmarkVecsProduct-2        	     100	  22024967 ns/op
BenchmarkVecsProductParallel-2	     100	  14118452 ns/op
ok  	_/home/simon/Dropbox/github/working/tensor3	12.933s
Tue 30 May 14:20:23 BST 2017
*/

