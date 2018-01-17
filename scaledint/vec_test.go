package tensor3

import "testing"
import "fmt"

func TestVecPrint(t *testing.T) {
	v := NewVector(1, 2, 3)
	if fmt.Sprint(v) != "{1 2 3}" {
		t.Error(v)
	}
}

func TestNewVector(t *testing.T) {
	v := *new(Vector)
	if fmt.Sprint(v) != "{0 0 0}" {
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
	v := NewVector(1, 2, 3)
	if v.Dot(NewVector(4, 5, 6)) != baseScale(32) {
		t.Error(v)
	}
}

func TestVecAdd(t *testing.T) {
	v := NewVector(1, 2, 3)
	v.Add(NewVector(4, 5, 6))
	if v != (NewVector(5, 7, 9)) {
		t.Error(v)
	}
}

func TestVecSubtract(t *testing.T) {
	v := NewVector(1, 2, 3)
	v.Subtract(NewVector(4, 5, 6))
	if v != (NewVector(-3, -3, -3)) {
		t.Error(v)
	}
}

func TestVecCross(t *testing.T) {
	v := NewVector(1, 2, 3)
	v.Cross(NewVector(4, 5, 6))
	if v != (NewVector(-3, 6, -3)) {
		t.Error(v)
	}
}

func TestVecLengthLength(t *testing.T) {
	v := NewVector(1, 2, 3)
	if v.LengthLength() != baseScale(14) {
		t.Error(v.LengthLength())
	}
}

func TestVecProduct(t *testing.T) {
	v := NewVector(1, 2, 3)
	v.Product(Identity)
	if v != (NewVector(1, 2, 3)) {
		t.Error(v)
	}
}

func TestVecProductT(t *testing.T) {
	v := NewVector(1, 2, 3)
	v.ProductT(Identity)
	if v != (NewVector(1, 2, 3)) {
		t.Error(v)
	}
}

func TestVecProject(t *testing.T) {
	v := NewVector(1, 2, 3)
	v.Project(NewVector(-2, 1, -1))
	if v != (NewVector(-2, 2, -3)) {
		t.Error(v)
	}
}

func TestVecLongestAxis(t *testing.T) {
	v := NewVector(1, 2, 3)
	if v.LongestAxis() != ZAxisIndex {
		t.Error(v.LongestAxis())
	}
	v.Subtract(NewVector(0, 0, 3))
	if v.LongestAxis() != YAxisIndex {
		t.Error(v.LongestAxis())
	}
	v.Subtract(NewVector(0, 2, 0))
	if v.LongestAxis() != XAxisIndex {
		t.Error(v.LongestAxis())
	}
}

func TestVecShortestAxis(t *testing.T) {
	v := NewVector(1, 2, 3)
	if v.ShortestAxis() != XAxisIndex {
		t.Error(v.ShortestAxis())
	}
	v.Subtract(NewVector(-1, 0, 3))
	if v.ShortestAxis() != ZAxisIndex {
		t.Error(v.ShortestAxis())
	}
	v.Subtract(NewVector(-1, 0, -3))
	if v.ShortestAxis() != YAxisIndex {
		t.Error(v.ShortestAxis())
	}
}

func TestVecMax(t *testing.T) {
	v := NewVector(1, 2, 3)
	v.Max(NewVector(-2, 1, -1))
	if v != (NewVector(1, 2, 3)) {
		t.Error(v)
	}
}

func TestVecMin(t *testing.T) {
	v := NewVector(1, 2, 3)
	v.Min(NewVector(-2, 1, -1))
	if v != (NewVector(-2, 1, -1)) {
		t.Error(v)
	}
}

func TestVecMid(t *testing.T) {
	v := NewVector(0, 5, 3)
	v.Mid(NewVector(-2, 1, -1))
	if v != (NewVector(-1, 3, 1)) {
		t.Error(v)
	}
}

func TestVecInterpolate(t *testing.T) {
	v := NewVector(1, 2, 3)
	v.Interpolate(NewVector(-2, 1, -1), 2)
	if v != (NewVector(4, 3, 7)) {
		t.Error(v)
	}
}

func TestVecApplyRunning(t *testing.T) {
	v := NewVector(1, 2, 3)
	v.Aggregate(Vectors{NewVector(1, 2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9)}, (*Vector).Add)
	if fmt.Sprint(v) != "{13 17 21}" {
		t.Error(v)
	}
}

func TestVecApplyForAll(t *testing.T) {
	Parallel = true
	defer func() {
		Parallel = false
	}()
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 2
	v := NewVector(1, 2, 3)
	v.ForAll(Vectors{NewVector(1, 2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9)}, (*Vector).Add)
	if fmt.Sprint(v) != "{13 17 21}" {
		t.Error(v)
	}
}
