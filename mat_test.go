package tensor3

import "testing"
import "fmt"

func TestMatrixIdentityPrint(t *testing.T) {
	if fmt.Sprint(Identity) != "{{1 0 0} {0 1 0} {0 0 1}}" {
		t.Error(Identity)
	}
}

func TestMatrixNew(t *testing.T) {
	v := *new(Matrix)
	if fmt.Sprint(v) != "{{0 0 0} {0 0 0} {0 0 0}}" {
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
	m := *new(Matrix)
	m.Add(NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9))
	if fmt.Sprint(m) != "{{1 2 3} {4 5 6} {7 8 9}}" {
		t.Error(m)
	}
}

func TestMatrixDeterminant(t *testing.T) {
	m := NewMatrix(2, 2, 3, 4, 5, 6, 7, 8, 9)
	if fmt.Sprint(m.Determinant()) != "-3" {
		t.Error(m.Determinant())
	}
}

func TestMatrixInvert(t *testing.T) {
	m := NewMatrix(2, 2, 3, 4, 5, 6, 7, 8, 9)
	var m2 Matrix
	m2.Set(m)
	m2.Invert()
	m.Product(m2)
	if m != Identity {
		t.Error(m)
	}
}

func TestMatrixProduct(t *testing.T) {
	m := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m.Product(NewMatrix(9, 8, 7, 6, 5, 4, 3, 2, 1))
	if fmt.Sprint(m) != "{{30 24 18} {84 69 54} {138 114 90}}" {
		t.Error(m)
	}
}

func TestMatrixProductDet(t *testing.T) {
	m1 := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m2 := NewMatrix(9, 8, 7, 6, 5, 4, 3, 2, 1)
	d1 := m1.Determinant()
	d2 := m2.Determinant()
	m1.Product(m2)
	dp := m1.Determinant()
	if d1*d2 != dp {
		t.Error("The determinant of a matrix product of square matrices did NOT equal the product of their determinants:")
	}
}

func TestMatrixInvertDet(t *testing.T) {
	m := NewMatrix(1, 2, 3, 4, -5, 6, 7, 8, 9)
	d := m.Determinant()
	m.Invert()
	di := m.Determinant()
	if float32(d*di) != 1 {
		//fmt.Println(d,di,d*di)
		t.Error("The determinant of the inverse of an invertible matrix should be the inverse of the determinant.")
	}
}

func TestMatrixProductIndentity(t *testing.T) {
	m := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m.Product(Identity)
	if fmt.Sprint(m) != "{{1 2 3} {4 5 6} {7 8 9}}" {
		t.Error(m)
	}
}

func TestMatrixMultiply(t *testing.T) {
	m := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m.Multiply(2)
	if fmt.Sprint(m) != "{{2 4 6} {8 10 12} {14 16 18}}" {
		t.Error(m)
	}
}

func TestMatrixAggregate(t *testing.T) {
	m := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m.Aggregate(Matrices{NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9), NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)}, (*Matrix).Add)
	if fmt.Sprint(m) != "{{3 6 9} {12 15 18} {21 24 27}}" {
		t.Error(m)
	}
}

func TestMatrixApplyXAdd(t *testing.T) {
	m := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m.applyX((*Vector).Add, NewVector(1, 2, 3))
	if fmt.Sprint(m) != "{{2 4 6} {4 5 6} {7 8 9}}" {
		t.Error(m)
	}
}

func TestMatrixApplyZAdd(t *testing.T) {
	m := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m.applyZ((*Vector).Add, NewVector(1, 2, 3))
	if fmt.Sprint(m) != "{{1 2 3} {4 5 6} {8 10 12}}" {
		t.Error(m)
	}
}

func TestMatrixApplyComponentWiseAxesAdd(t *testing.T) {
	m := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m.ApplyToComponentsByAxes((*Vector).Add)
	if fmt.Sprint(m) != "{{2 2 3} {4 6 6} {7 8 10}}" {
		t.Error(m)
	}
}

func TestMatrixApplyComponentWiseAxesCross(t *testing.T) {
	m := Matrix{zAxis, xAxis, yAxis}
	m.ApplyToComponentsByAxes((*Vector).Cross)
	m2 := Matrix{yAxis, zAxis, xAxis}
	if m != m2 {
		t.Error(m)
	}
}

func BenchmarkMatrixProduct(b *testing.B) {
	b.StopTimer()
	m := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m2 := NewMatrix(9, 8, 7, 6, 5, 4, 3, 2, 1)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Product(m2)
	}
}

/*  Hal3 Sun 28 May 20:13:58 BST 2017 go version go1.6.2 linux/amd64
PASS
BenchmarkMatrixProduct-2      	50000000	        31.5 ns/op
BenchmarkMatsProduct-2        	      30	  41442969 ns/op
BenchmarkMatsProductParallel-2	      30	  41318688 ns/op
ok  	_/home/simon/Dropbox/github/working/tensor3	4.641s
Sun 28 May 20:14:04 BST 2017
*/
/*  Hal3 Sun 28 May 20:14:18 BST 2017  go version go1.8.3 linux/amd64

BenchmarkMatrixProduct-2         	50000000	        27.4 ns/op
BenchmarkMatsProduct-2           	      30	  44000097 ns/op
BenchmarkMatsProductParallel-2   	      30	  41316237 ns/op
PASS
ok  	_/home/simon/Dropbox/github/working/tensor3	4.383s
Sun 28 May 20:14:29 BST 2017
*/
