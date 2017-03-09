package tensor3

import "testing"
import "fmt"

func TestMatrixIdentityPrint(t *testing.T) {
	if fmt.Sprint(Identity) != "{{1 0 0} {0 1 0} {0 0 1}}" {
		t.Error("identity")
	}
}
func TestMatrixNew(t *testing.T) {
	v := new(Matrix)
	if fmt.Sprint(v) != "&{{0 0 0} {0 0 0} {0 0 0}}" {
		t.Error(v)
	}
}

func TestMatrixPrint(t *testing.T) {
	m := NewMatrix(1, 2, 3)
	if fmt.Sprint(m) != "{{1 2 3} {0 0 0} {0 0 0}}" {
		t.Error(m)
	}
}
func TestMatrixAdd(t *testing.T) {
	m := new(Matrix)
	m.Add(Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}})
	if fmt.Sprint(m) != "&{{1 2 3} {4 5 6} {7 8 9}}" {
		t.Error(m)
	}
}

func TestMatrixProduct(t *testing.T) {
	m := Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	m.Product(Matrix{Vector{9, 8, 7}, Vector{6, 5, 4}, Vector{3, 2, 1}})
	if fmt.Sprint(m) != "{{30 24 18} {84 69 54} {138 114 90}}" {
		t.Error(m)
	}
}

func TestMatrixProductT(t *testing.T) {
	m := Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	m2 := Matrix{Vector{9, 8, 7}, Vector{6, 5, 4}, Vector{3, 2, 1}}
	m2.Transpose()
	m.ProductT(m2)
	if fmt.Sprint(m) != "{{30 24 18} {84 69 54} {138 114 90}}" {
		t.Error(m)
	}
}

func TestMatrixProductIndentity(t *testing.T) {
	m := Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	m.Product(Identity)
	if fmt.Sprint(m) != "{{1 2 3} {4 5 6} {7 8 9}}" {
		t.Error(m)
	}
}

func TestMatrixMultiply(t *testing.T) {
	m := Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	m.Multiply(2)
	if fmt.Sprint(m) != "{{2 4 6} {8 10 12} {14 16 18}}" {
		t.Error(m)
	}
}

func TestMatrixReduce(t *testing.T) {
	m := Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	m.Reduce(Matrices{Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}, Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}}, (*Matrix).Add)
	if fmt.Sprint(m) != "{{3 6 9} {12 15 18} {21 24 27}}" {
		t.Error(m)
	}
}

func TestMatrixApplyXAdd(t *testing.T) {
	m := Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	m.applyX((*Vector).Add, Vector{1, 2, 3})
	if fmt.Sprint(m) != "{{2 4 6} {4 5 6} {7 8 9}}" {
		t.Error(m)
	}
}

func TestMatrixApplyZAdd(t *testing.T) {
	m := Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	m.applyZ((*Vector).Add, Vector{1, 2, 3})
	if fmt.Sprint(m) != "{{1 2 3} {4 5 6} {8 10 12}}" {
		t.Error(m)
	}
}

func BenchmarkMatrixProduct(b *testing.B) {
	b.StopTimer()
	m := Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	m2 := Matrix{Vector{9, 8, 7}, Vector{6, 5, 4}, Vector{3, 2, 1}}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Product(m2)
	}
}
