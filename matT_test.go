package tensor3

import "testing"
import "fmt"

func TestMatrixProductT(t *testing.T) {
	m := Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	m2 := Matrix{Vector{9, 6, 3}, Vector{8, 5, 2}, Vector{7, 4, 1}}
	m2.Transpose()
	m.ProductT(m2)
	if fmt.Sprint(m) != "{{30 24 18} {84 69 54} {138 114 90}}" {
		//	if fmt.Sprint(m) != "{{46 118 190} {28 73 118} {10 28 46}}" {
		t.Error(m)
	}
}

func TestMatrixTProductT(t *testing.T) {
	m := Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	m2 := Matrix{Vector{9, 6, 3}, Vector{8, 5, 2}, Vector{7, 4, 1}}
	m2.Transpose()
	m.TProductT(m2)
	m.Transpose()
	if fmt.Sprint(m) != "{{30 24 18} {84 69 54} {138 114 90}}" {
		//	if fmt.Sprint(m) != "{{46 118 190} {28 73 118} {10 28 46}}" {
		t.Error(m)
	}
}
