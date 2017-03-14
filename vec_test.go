package tensor3

import "testing"
import "fmt"

func TestVecPrint(t *testing.T) {
	v := Vector{1, 2, 3}
	if fmt.Sprint(v) != "{1 2 3}" {
		t.Error(v)
	}
}

func TestNewVector(t *testing.T) {
	v := new(Vector)
	if fmt.Sprint(v) != "&{0 0 0}" {
		t.Error(v)
	}
}

func TestNewVectorPrint(t *testing.T) {
	v := NewVector(1, 2, 3)
	if fmt.Sprint(v) != "{1 2 3}" {
		t.Error(v)
	}
}

func TestVecDot(t *testing.T) {
	v := Vector{1, 2, 3}
	if v.Dot(Vector{4, 5, 6}) != 32.0 {
		t.Error(v)
	}
}

func TestVecAdd(t *testing.T) {
	v := Vector{1, 2, 3}
	v.Add(Vector{4, 5, 6})
	if v != (Vector{5, 7, 9}) {
		t.Error(v)
	}
}

func TestVecSubtract(t *testing.T) {
	v := Vector{1, 2, 3}
	v.Subtract(Vector{4, 5, 6})
	if v != (Vector{-3, -3, -3}) {
		t.Error(v)
	}
}

func TestVecCross(t *testing.T) {
	v := Vector{1, 2, 3}
	v.Cross(Vector{4, 5, 6})
	if v != (Vector{-3, 6, -3}) {
		t.Error(v)
	}
}

func TestVecLengthLength(t *testing.T) {
	v := Vector{1, 2, 3}
	if v.LengthLength() != 14 {
		t.Error(v)
	}
}

func TestVecProduct(t *testing.T) {
	v := Vector{1, 2, 3}
	v.Product(Identity)
	if v != (Vector{1, 2, 3}) {
		t.Error(v)
	}
}

func TestVecProductT(t *testing.T) {
	v := Vector{1, 2, 3}
	v.ProductT(Identity)
	if v != (Vector{1, 2, 3}) {
		t.Error(v)
	}
}

func TestVecProject(t *testing.T) {
	v := Vector{1, 2, 3}
	v.Project(Vector{-2, 1, -1})
	if v != (Vector{-2, 2, -3}) {
		t.Error(v)
	}
}

func TestVecLongestAxis(t *testing.T) {
	v := Vector{1, 2, 3}
	if v.LongestAxis() != zAxisIndex {
		t.Error(v.LongestAxis())
	}
	v.Subtract( Vector{0, 0, 3})
	if v.LongestAxis() != yAxisIndex {
		t.Error(v.LongestAxis())
	}
	v.Subtract( Vector{4, 0, 0})
	if v.LongestAxis() != xAxisIndex {
		t.Error(v.LongestAxis())
	}
}


func TestVecMax(t *testing.T) {
	v := Vector{1, 2, 3}
	v.Max(Vector{-2, 1, -1})
	if v != (Vector{1, 2, 3}) {
		t.Error(v)
	}
}

func TestVecMin(t *testing.T) {
	v := Vector{1, 2, 3}
	v.Min(Vector{-2, 1, -1})
	if v != (Vector{-2, 1, -1}) {
		t.Error(v)
	}
}

func TestVecMid(t *testing.T) {
	v := Vector{1, 2, 3}
	v.Mid(Vector{-2, 1, -1})
	if v != (Vector{-0.5, 1.5, 1}) {
		t.Error(v)
	}
}

func TestVecInterpolate(t *testing.T) {
	v := Vector{1, 2, 3}
	v.Interpolate(Vector{-2, 1, -1}, 0.5)
	if v != (Vector{-0.5, 1.5, 1}) {
		t.Error(v)
	}
}

func TestVecApplyRunning(t *testing.T) {
	v := Vector{1, 2, 3}
	v.Reduce(Vectors{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}, (*Vector).Add)
	if fmt.Sprint(v) != "{13 17 21}" {
		t.Error(v)
	}
}
