package tensor3

import "testing"
import "fmt"

func TestVecRefsNewFromIndexes(t *testing.T) {
	vr := NewVectorRefsFromIndexes(NewVectors(1, 2, 3, 4, 5, 6, 7), 1, 3, 2)
	if fmt.Sprint(*vr[2]) != "{4 5 6}" {
		t.Error(fmt.Sprint(*vr[2]))
	}
}

func TestVecRefsNewFromEmplyIndexes(t *testing.T) {
	vr := NewVectorRefsFromIndexes(NewVectors(1, 2, 3, 4, 5, 6, 7))
	if fmt.Sprint(*vr[2]) != "{7 0 0}" {
		t.Error(fmt.Sprint(*vr[2]))
	}
}

func TestVecRefsDereferenceNew(t *testing.T) {
	vr := NewVectorRefsFromIndexes(NewVectors(1, 2, 3, 4, 5, 6, 7), 1, 3, 2)
	if fmt.Sprint(vr.Dereference()) != "[{1 2 3} {7 0 0} {4 5 6}]" {
		t.Error(fmt.Sprint(vr.Dereference()))
	}
}

func TestVecRefsDereference(t *testing.T) {
	vr := NewVectorRefsFromIndexes(NewVectors(1, 2, 3, 4, 5, 6, 7), 1, 3, 2)
	v := make(Vectors, 2)
	v.Reference(vr)
	if fmt.Sprint(v) != "[{1 2 3} {7 0 0}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsFromVectorRefs(t *testing.T) {
	vs := NewVectors(1, 2, 3, 4, 5, 6, 7)
	vr1 := NewVectorRefsFromIndexes(vs, 1)
	vr2 := NewVectorRefsFromIndexes(vs, 3)
	vr := NewVectorsFromVectorRefs(vr1, vr2)
	if fmt.Sprint(vr) != "[{1 2 3} {7 0 0}]" {
		t.Error(fmt.Sprint(vr))
	}
	if fmt.Sprint(vr1.Dereference()) != "[{1 2 3}]" {
		t.Error(fmt.Sprint(vr1.Dereference()))
	}
	if fmt.Sprint(vr2.Dereference()) != "[{7 0 0}]" {
		t.Error(fmt.Sprint(vr2.Dereference()))
	}
}

func TestVecsIndexesFromVectorRefs(t *testing.T) {
	vs := NewVectors(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	vr1 := NewVectorRefsFromIndexes(vs, 3, 5, 1)
	vr2 := NewVectorRefsFromIndexes(vs, 2, 5)
	vr := NewVectorsFromVectorRefs(vr2, vr1)
	is1 := vr1.Indexes(vr)
	is2 := vr2.Indexes(vr)
	if fmt.Sprint(is1, is2, vr) != "[3 2 4] [1 2] [{4 5 6} {13 14 15} {7 8 9} {1 2 3}]" {
		t.Error(fmt.Sprint(is1, is2, vr))
	}

}

func TestVecRefsPrint(t *testing.T) {
	vr := VectorRefs{&Vector{1 * scale, 2 * scale, 3 * scale}}
	if fmt.Sprint(*vr[0]) != "{1 2 3}" {
		t.Error(fmt.Sprint(*vr[0]))
	}
}

func TestVecRefsSum(t *testing.T) {
	vs := VectorRefs{&Vector{7 * scale, 8 * scale, 9 * scale}, &Vector{7 * scale, 8 * scale, 9 * scale}, &Vector{7 * scale, 8 * scale, 9 * scale}}
	if fmt.Sprint(vs.Sum()) != "{21 24 27}" {
		t.Error(vs.Sum())
	}
}

func TestVecRefsAddRefs(t *testing.T) {
	vs := VectorRefs{&Vector{1 * scale, 2 * scale, 3 * scale}, &Vector{4 * scale, 5 * scale, 6 * scale}, &Vector{7 * scale, 8 * scale, 9 * scale}}
	vs2 := VectorRefs{&Vector{9 * scale, 8 * scale, 7 * scale}, &Vector{6 * scale, 5 * scale, 4 * scale}, &Vector{3 * scale, 2 * scale, 1 * scale}}
	vs.AddAllRefs(vs2)
	if fmt.Sprint(*vs[0], *vs[1], *vs[2]) != "{10 10 10} {10 10 10} {10 10 10}" {
		t.Error(fmt.Sprint(*vs[0], *vs[1], *vs[2]))
	}
}

func TestVecRefsCrossRefs(t *testing.T) {
	vs := VectorRefs{&Vector{1 * scale, 2 * scale, 3 * scale}, &Vector{4 * scale, 5 * scale, 6 * scale}, &Vector{7 * scale, 8 * scale, 9 * scale}}
	vs2 := VectorRefs{&Vector{9 * scale, 8 * scale, 7 * scale}, &Vector{6 * scale, 5 * scale, 4 * scale}, &Vector{3 * scale, 2 * scale, 1 * scale}}
	vs.CrossAllRefs(vs2)
	if fmt.Sprint(*vs[0], *vs[1], *vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}" {
		t.Error(fmt.Sprint(*vs[0], *vs[1], *vs[2]))
	}
}

func TestVecsAddVecRefs(t *testing.T) {
	vs := Vectors{NewVector(1, 2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9)}
	vs2 := VectorRefs{&Vector{9 * scale, 8 * scale, 7 * scale}, &Vector{6 * scale, 5 * scale, 4 * scale}, &Vector{3 * scale, 2 * scale, 1 * scale}}
	vs.AddAllRefs(vs2)
	if fmt.Sprint(vs[0], vs[1], vs[2]) != "{10 10 10} {10 10 10} {10 10 10}" {
		t.Error(fmt.Sprint(vs[0], vs[1], vs[2]))
	}
}

func TestVecsCrossVecRefs(t *testing.T) {
	vs := Vectors{NewVector(1, 2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9)}
	vs2 := VectorRefs{&Vector{9 * scale, 8 * scale, 7 * scale}, &Vector{6 * scale, 5 * scale, 4 * scale}, &Vector{3 * scale, 2 * scale, 1 * scale}}
	vs.CrossAllRefs(vs2)
	if fmt.Sprint(vs[0], vs[1], vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}" {
		t.Error(fmt.Sprint(vs[0], vs[1], vs[2]))
	}
}
func TestVecsRefsAddVecs(t *testing.T) {
	vs := VectorRefs{&Vector{1 * scale, 2 * scale, 3 * scale}, &Vector{4 * scale, 5 * scale, 6 * scale}, &Vector{7 * scale, 8 * scale, 9 * scale}}
	vs2 := Vectors{NewVector(9, 8, 7), NewVector(6, 5, 4), NewVector(3, 2, 1)}
	vs.AddAll(vs2)
	if fmt.Sprint(*vs[0], *vs[1], *vs[2]) != "{10 10 10} {10 10 10} {10 10 10}" {
		t.Error(fmt.Sprint(*vs[0], *vs[1], *vs[2]))
	}
}

func TestVecRefsCrossVecs(t *testing.T) {
	vs := VectorRefs{&Vector{1 * scale, 2 * scale, 3 * scale}, &Vector{4 * scale, 5 * scale, 6 * scale}, &Vector{7 * scale, 8 * scale, 9 * scale}}
	vs2 := Vectors{NewVector(9, 8, 7), NewVector(6, 5, 4), NewVector(3, 2, 1)}
	vs.CrossAll(vs2)
	if fmt.Sprint(*vs[0], *vs[1], *vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}" {
		t.Error(fmt.Sprint(*vs[0], *vs[1], *vs[2]))
	}
}

func TestVecsRefsAddVecRefs(t *testing.T) {
	vs := VectorRefs{&Vector{1 * scale, 2 * scale, 3 * scale}, &Vector{4 * scale, 5 * scale, 6 * scale}, &Vector{7 * scale, 8 * scale, 9 * scale}}
	vs2 := VectorRefs{&Vector{9 * scale, 8 * scale, 7 * scale}, &Vector{6 * scale, 5 * scale, 4 * scale}, &Vector{3 * scale, 2 * scale, 1 * scale}}
	vs.AddAllRefs(vs2)
	if fmt.Sprint(*vs[0], *vs[1], *vs[2]) != "{10 10 10} {10 10 10} {10 10 10}" {
		t.Error(fmt.Sprint(*vs[0], *vs[1], *vs[2]))
	}
}

func TestVecRefsCrossVecRefs(t *testing.T) {
	vs := VectorRefs{&Vector{1 * scale, 2 * scale, 3 * scale}, &Vector{4 * scale, 5 * scale, 6 * scale}, &Vector{7 * scale, 8 * scale, 9 * scale}}
	vs2 := VectorRefs{&Vector{9 * scale, 8 * scale, 7 * scale}, &Vector{6 * scale, 5 * scale, 4 * scale}, &Vector{3 * scale, 2 * scale, 1 * scale}}
	vs.CrossAllRefs(vs2)
	if fmt.Sprint(*vs[0], *vs[1], *vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}" {
		t.Error(fmt.Sprint(*vs[0], *vs[1], *vs[2]))
	}
}

func TestVecRefsSelect(t *testing.T) {
	vs := VectorRefs{&Vector{1 * scale, 2 * scale, 3 * scale}, &Vector{4 * scale, 5 * scale, 6 * scale}, &Vector{7 * scale, 8 * scale, 9 * scale}}
	vs2:= vs.Select(
		func(vr *Vector)bool{
			return vr.y==5*scale
		}
	)
	if fmt.Sprint(vs2) != "[{4 5 6}]" {
		t.Error(fmt.Sprint(vs2))
	}
}

func BenchmarkVecRefsProduct(b *testing.B) {
	b.StopTimer()
	vrs := make(VectorRefs, 100000)
	for i := range vrs {
		vrs[i] = &Vector{1 * scale, 2 * scale, 3 * scale}
	}
	m := Matrix{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vrs.Product(m)
	}

}

func BenchmarkVecRefsProductParallel(b *testing.B) {
	b.StopTimer()
	vrs := make(VectorRefs, 100000)
	for i := range vrs {
		vrs[i] = &Vector{1 * scale, 2 * scale, 3 * scale}
	}
	m := Matrix{}
	Parallel = true
	defer func() {
		Parallel = false
	}()
	Hints.ChunkSizeFixed = true
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vrs.Product(m)
	}
}

