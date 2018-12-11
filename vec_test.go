package tensor3

import "testing"
import "fmt"

func TestVecPrint(t *testing.T) {
	v := *New(1, 2, 3)
	if fmt.Sprint(v) != "{1 2 3}" {
		t.Error(v)
	}
}

func TestNew(t *testing.T) {
	v := *new(Vector)
	if fmt.Sprint(v) != "{0 0 0}" {
		t.Error(v)
	}
}

func TestNewVectorPrint(t *testing.T) {
	v := *New(1, 2, 3)
	if fmt.Sprint(v) != "{1 2 3}" {
		t.Error(v)
	}
}

func TestVecDot(t *testing.T) {
	v := *New(1, 2, 3)
	if fmt.Sprint(v.Dot(*New(4, 5, 6))) != "32" {
		t.Error(v.Dot(*New(4, 5, 6)))
	}
}

func TestVecAdd(t *testing.T) {
	v := New(1, 2, 3)
	v.Add(*New(4, 5, 6))
	if *v != (*New(5, 7, 9)) {
		t.Error(v)
	}
}

func TestVecSubtract(t *testing.T) {
	v := New(1, 2, 3)
	v.Subtract(*New(4, 5, 6))
	if *v != (*New(-3, -3, -3)) {
		t.Error(v)
	}
}

func TestVecCross(t *testing.T) {
	v := New(1, 2, 3)
	v.Cross(*New(4, 5, 6))
	if *v != (*New(-3, 6, -3)) {
		t.Error(v)
	}
}

func TestVecLengthLength(t *testing.T) {
	v := New(1, 2, 3)
	if fmt.Sprint(v.LengthLength()) != "14" {
		t.Error(v.LengthLength())
	}
}

func TestVecDistDist(t *testing.T) {
	v := New(3, 2, 4)
	if fmt.Sprint(v.DistDist(*New(-3, 2, -4))) != "100" {
		t.Error(v.DistDist(*New(-3, 2, -4)))
	}
}

func TestVecDistDistSame(t *testing.T) {
	v := New(1, 2, 3)
	if v.DistDist(*New(1, 2, 3)) != 0 {
		t.Error(v.DistDist(*New(1, 2, 3)))
	}
}

func TestVecProduct(t *testing.T) {
	v := New(1, 2, 3)
	v.Product(Identity)
	if *v != (*New(1, 2, 3)) {
		t.Error(v)
	}
}

func TestVecProductT(t *testing.T) {
	v := New(1, 2, 3)
	v.ProductT(Identity)
	if *v != (*New(1, 2, 3)) {
		t.Error(v)
	}
}

func TestVecProject(t *testing.T) {
	v := New(1, 2, 3)
	v.Project(*New(-2, 1, -1))
	if *v != (*New(-2, 2, -3)) {
		t.Error(v)
	}
}

func TestVecProjectXYZ(t *testing.T) {
	v := New(1, 2, 3)
	x, y := v.ProjectX()
	if x != 2 || y != 3 {
		t.Error()
	}
	x, y = v.ProjectY()
	if x != 3 || y != 1 {
		t.Error()
	}
	x, y = v.ProjectZ()
	if x != 1 || y != 2 {
		t.Error()
	}
}

func TestVecLongestAxis(t *testing.T) {
	v := New(1, 2, 3)
	if v.LongestAxis() != ZAxisIndex {
		t.Error(v.LongestAxis())
	}
	v.Subtract(*New(0, 0, 3))
	if v.LongestAxis() != YAxisIndex {
		t.Error(v.LongestAxis())
	}
	v.Subtract(*New(0, 2, 0))
	if v.LongestAxis() != XAxisIndex {
		t.Error(v.LongestAxis())
	}
}

func TestVecShortestAxis(t *testing.T) {
	v := New(1, 2, 3)
	if v.ShortestAxis() != XAxisIndex {
		t.Error(v.ShortestAxis())
	}
	v.Subtract(*New(-1, 0, 3))
	if v.ShortestAxis() != ZAxisIndex {
		t.Error(v.ShortestAxis())
	}
	v.Subtract(*New(-1, 0, -3))
	if v.ShortestAxis() != YAxisIndex {
		t.Error(v.ShortestAxis())
	}
}

func TestVecMax(t *testing.T) {
	v := New(1, 2, 3)
	v.Max(*New(-2, 1, -1))
	if *v != (*New(1, 2, 3)) {
		t.Error(v)
	}
}

func TestVecMin(t *testing.T) {
	v := New(1, 2, 3)
	v.Min(*New(-2, 1, -1))
	if *v != (*New(-2, 1, -1)) {
		t.Error(v)
	}
}

func TestVecMid(t *testing.T) {
	v := New(0, 5, 3)
	v.Mid(*New(-2, 1, -1))
	if *v != (*New(-1, 3, 1)) {
		t.Error(v)
	}
}

func TestVecInterpolate(t *testing.T) {
	v := New(1, 2, 3)
	v.Interpolate(*New(-2, 1, -1), .5)
	if *v != (*New(-.5, 1.5, 1)) {
		t.Error(v)
	}
}

func TestVecApplyRunning(t *testing.T) {
	v := New(1, 2, 3)
	v.Aggregate(Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9)}, (*Vector).Add)
	if fmt.Sprint(*v) != "{13 17 21}" {
		t.Error(v)
	}
}

func TestVecApplyForAll(t *testing.T) {
	Parallel = true
	defer func(d int) {
		Parallel = false
		Hints.DefaultChunkSize = d
	}(Hints.DefaultChunkSize)
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 2
	v := *New(1, 2, 3)
	v.ForAll(Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9)}, (*Vector).Add)
	if fmt.Sprint(v) != "{13 17 21}" {
		t.Error(v)
	}
}
