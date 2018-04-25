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
	v := NewVectors(1, 2, 3, 4, 5, 6, 7)
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
			t.Log("As expected:", r)
		} else {
			t.Error("Expected error not present.")
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
	defer func(d uint) {
		Parallel = false
		Hints.DefaultChunkSize = d
	}(Hints.DefaultChunkSize)
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
	defer func(d uint) {
		Parallel = false
		Hints.DefaultChunkSize = d
	}(Hints.DefaultChunkSize)
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
	defer func(d uint) {
		Parallel = false
		Hints.DefaultChunkSize = d
	}(Hints.DefaultChunkSize)
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 2
	v := Vectors{NewVector(1, 2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9), NewVector(10, 11, 12), NewVector(13, 14, 15), NewVector(16, 17, 18)}
	//	v=append(v,Vector{1, 2, 3})
	v.Add(NewVector(1, 1, 1))
	if fmt.Sprint(v) != "[{2 3 4} {5 6 7} {8 9 10} {11 12 13} {14 15 16} {17 18 19}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsAddVecs(t *testing.T) {
	vs := Vectors{NewVector(1, 2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9)}
	vs2 := Vectors{NewVector(9, 8, 7), NewVector(6, 5, 4), NewVector(3, 2, 1)}
	vs.AddAll(vs2)
	if fmt.Sprint(vs[0], vs[1], vs[2]) != "{10 10 10} {10 10 10} {10 10 10}" {
		t.Error(fmt.Sprint(vs[0], vs[1], vs[2]))
	}
}

func TestVecsCrossVecs(t *testing.T) {
	vs := Vectors{NewVector(1, 2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9)}
	vs2 := Vectors{NewVector(9, 8, 7), NewVector(6, 5, 4), NewVector(3, 2, 1)}
	vs.CrossAll(vs2)
	if fmt.Sprint(vs[0], vs[1], vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}" {
		t.Error(fmt.Sprint(vs[0], vs[1], vs[2]))
	}
}

func BenchmarkVecsSum(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 100000)
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
	vs := make(Vectors, 100000)
	for i := range vs {
		vs[i] = NewVector(1, 2, 3)
	}
	Parallel = true
	defer func() {
		Parallel = false
	}()
	Hints.ChunkSizeFixed = true
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Sum()
	}
}

func BenchmarkVecsCross(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 100000)
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
	vs := make(Vectors, 100000)
	for i := range vs {
		vs[i] = NewVector(1, 2, 3)
	}
	v := NewVector(9, 8, 7)
	Parallel = true
	defer func() {
		Parallel = false
	}()
	Hints.ChunkSizeFixed = true
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Cross(v)
	}

}

func BenchmarkVecsProduct(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 100000)
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
	vs := make(Vectors, 100000)
	for i := range vs {
		vs[i] = NewVector(1, 2, 3)
	}
	m := Matrix{}
	Parallel = true
	defer func() {
		Parallel = false
	}()
	Hints.ChunkSizeFixed = true
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Product(m)
	}

}

/*
tensor3.test -test.bench Vecs
As expected: runtime error: index out of range
goos: linux
goarch: arm
BenchmarkVecsSum-4               	     100	  12240145 ns/op
BenchmarkVecsSumParallel-4       	     100	  11940562 ns/op
BenchmarkVecsCross-4             	     200	   7850206 ns/op
BenchmarkVecsCrossParallel-4     	     500	   2793741 ns/op
BenchmarkVecsProduct-4           	     100	  11736753 ns/op
BenchmarkVecsProductParallel-4   	     500	   3939572 ns/op
PASS
*/

