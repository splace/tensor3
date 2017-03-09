package tensor3

import "testing"
import "fmt"

func TestMatsApplyComponentsXYZ(t *testing.T) {
	ms := Matrices{Matrix{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}}
	ms.ApplyComponentWiseVariac(Vector{1, 2, 3}, (*Vector).Add, (*Vector).Subtract, (*Vector).SetZ)
	if fmt.Sprint(ms) != "[{{2 4 6} {3 3 3} {7 8 3}}]" {
		t.Error(fmt.Sprint(ms))
	}
}
