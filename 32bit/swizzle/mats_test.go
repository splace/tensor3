package tensor3

import "testing"
import "fmt"

func TestMatsPrint(t *testing.T) {
	ms := Matrices{Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}}
	if fmt.Sprint(ms) != "[{{1 2 3} {4 5 6} {7 8 9}}]" {
		t.Error(fmt.Sprint(ms))
	}
}

func TestMatsProduct1(t *testing.T) {
	ms := Matrices{Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}}
	ms.Product(Matrix{Vector{9, 8, 7}, Vector{6, 5, 4}, Vector{3, 2, 1}})
	if fmt.Sprint(ms) != "[{{30 24 18} {84 69 54} {138 114 90}}]" {
		t.Error(fmt.Sprint(ms))
	}
}

func TestMatsSum(t *testing.T) {
	ms := Matrices{Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}, Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}, Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}}
	if fmt.Sprint(ms.Sum()) != "{{3 6 9} {12 15 18} {21 24 27}}" {
		t.Error(ms.Sum())
	}
}

func TestMatsApplyComponents(t *testing.T) {
	ms := Matrices{Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}}
	ms.ReduceComponentWise(Vector{1, 2, 3}, (*Vector).Add)
	if fmt.Sprint(ms) != "[{{2 4 6} {5 7 9} {8 10 12}}]" {
		t.Error(fmt.Sprint(ms))
	}
}

func TestMatsApplyComponentY(t *testing.T) {
	ms := Matrices{Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}}
	ms.ApplyComponentWiseVariac(Vector{1, 2, 3}, nil, (*Vector).Subtract)
	if fmt.Sprint(ms) != "[{{1 2 3} {3 3 3} {7 8 9}}]" {
		t.Error(fmt.Sprint(ms))
	}
}

func TestMatsApplyComponentsXY(t *testing.T) {
	ms := Matrices{Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}}
	ms.ApplyComponentWiseVariac(Vector{1, 2, 3}, (*Vector).Add, (*Vector).Subtract)
	if fmt.Sprint(ms) != "[{{2 4 6} {3 3 3} {7 8 9}}]" {
		t.Error(fmt.Sprint(ms))
	}
}

func BenchmarkMatsProduct(b *testing.B) {
	Parallel = false
	b.StopTimer()
	ms := make(Matrices, 1000000)
	for i := range ms {
		ms[i] = Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	}
	m := Matrix{Vector{9, 8, 7}, Vector{6, 5, 4}, Vector{3, 2, 1}}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ms.Product(m)
	}
}

func BenchmarkMatsProductParallel(b *testing.B) {
	b.StopTimer()
	ms := make(Matrices, 1000000)
	for i := range ms {
		ms[i] = Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	}
	m := Matrix{Vector{9, 8, 7}, Vector{6, 5, 4}, Vector{3, 2, 1}}
	Parallel = true
	Hints.ChunkSizeFixed = true
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ms.Product(m)
	}

}


/*  Hal3 Thu 9 Mar 16:55:05 GMT 2017 go version go1.6.2 linux/amd64
PASS
BenchmarkMatsProduct-2        	      30	  44015188 ns/op
BenchmarkMatsProductParallel-2	      30	  40866053 ns/op
ok  	2.989s
Thu 9 Mar 16:55:10 GMT 2017
*/
/*  Hal3 Thu 9 Mar 16:56:07 GMT 2017  go version go1.8 linux/amd64

BenchmarkMatsProduct-2           	      30	  43786364 ns/op
BenchmarkMatsProductParallel-2   	      30	  43681382 ns/op
PASS
ok  	3.110s
Thu 9 Mar 16:56:11 GMT 2017
*/

