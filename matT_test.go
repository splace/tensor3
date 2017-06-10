package tensor3

import "testing"
import "fmt"

func TestMatrixProductT(t *testing.T) {
	m := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m2 := NewMatrix(9, 8, 7, 6, 5, 4, 3, 2, 1)
	m2.Transpose()
	m.ProductT(m2)
	if fmt.Sprint(m) != "{{30 24 18} {84 69 54} {138 114 90}}" {
		t.Error(m)
	}
}

func TestMatrixTProduct(t *testing.T) {
	m := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m2 := NewMatrix(9, 8, 7, 6, 5, 4, 3, 2, 1)
	m.TProduct(m2)
	m.Transpose()
	if fmt.Sprint(m) != "{{30 24 18} {84 69 54} {138 114 90}}" {
		t.Error(m)
	}
}

func TestMatrixTProductT(t *testing.T) {
	m := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
	m2 := NewMatrix(9, 8, 7, 6, 5, 4, 3, 2, 1)
	m2.Transpose()
	m.TProductT(m2)
	m.Transpose()
	if fmt.Sprint(m) != "{{30 24 18} {84 69 54} {138 114 90}}" {
		t.Error(m)
	}
}
