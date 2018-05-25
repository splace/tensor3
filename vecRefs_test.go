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
	vr := VectorRefs{New(1,2,3)}
	if fmt.Sprint(*vr[0]) != "{1 2 3}" {
		t.Error(fmt.Sprint(*vr[0]))
	}
}

func TestVecRefsSum(t *testing.T) {
	vs := VectorRefs{New(7,8,9), New(7,8,9), New(7,8,9)}
	if fmt.Sprint(vs.Sum()) != "{21 24 27}" {
		t.Error(vs.Sum())
	}
}

func TestVecRefsAddRefs(t *testing.T) {
	vs := VectorRefs{New(1,2,3), New(4,5,6), New(7,8,9)}
	vs2 := VectorRefs{New(9,8,7), New(6,5,4), New(3,2,1)}
	vs.AddAllRefs(vs2)
	if fmt.Sprint(*vs[0], *vs[1], *vs[2]) != "{10 10 10} {10 10 10} {10 10 10}" {
		t.Error(fmt.Sprint(*vs[0], *vs[1], *vs[2]))
	}
}

func TestVecRefsCrossRefs(t *testing.T) {
	vs := VectorRefs{New(1,2,3), New(4,5,6), New(7,8,9)}
	vs2 := VectorRefs{New(9,8,7), New(6,5,4), New(3,2,1)}
	vs.CrossAllRefs(vs2)
	if fmt.Sprint(*vs[0], *vs[1], *vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}" {
		t.Error(fmt.Sprint(*vs[0], *vs[1], *vs[2]))
	}
}

func TestVecsAddVecRefs(t *testing.T) {
	vs := Vectors{NewVector(1, 2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9)}
	vs2 := VectorRefs{New(9,8,7), New(6,5,4), New(3,2,1)}
	vs.AddAllRefs(vs2)
	if fmt.Sprint(vs[0], vs[1], vs[2]) != "{10 10 10} {10 10 10} {10 10 10}" {
		t.Error(fmt.Sprint(vs[0], vs[1], vs[2]))
	}
}

func TestVecsCrossVecRefs(t *testing.T) {
	vs := Vectors{NewVector(1, 2, 3), NewVector(4, 5, 6), NewVector(7, 8, 9)}
	vs2 := VectorRefs{New(9,8,7), New(6,5,4), New(3,2,1)}
	vs.CrossAllRefs(vs2)
	if fmt.Sprint(vs[0], vs[1], vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}" {
		t.Error(fmt.Sprint(vs[0], vs[1], vs[2]))
	}
}
func TestVecsRefsAddVecs(t *testing.T) {
	vs := VectorRefs{New(1,2,3), New(4,5,6), New(7,8,9)}
	vs2 := Vectors{NewVector(9, 8, 7), NewVector(6, 5, 4), NewVector(3, 2, 1)}
	vs.AddAll(vs2)
	if fmt.Sprint(*vs[0], *vs[1], *vs[2]) != "{10 10 10} {10 10 10} {10 10 10}" {
		t.Error(fmt.Sprint(*vs[0], *vs[1], *vs[2]))
	}
}

func TestVecRefsCrossVecs(t *testing.T) {
	vs := VectorRefs{New(1,2,3), New(4,5,6), New(7,8,9)}
	vs2 := Vectors{NewVector(9, 8, 7), NewVector(6, 5, 4), NewVector(3, 2, 1)}
	vs.CrossAll(vs2)
	if fmt.Sprint(*vs[0], *vs[1], *vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}" {
		t.Error(fmt.Sprint(*vs[0], *vs[1], *vs[2]))
	}
}

func TestVecsRefsAddVecRefs(t *testing.T) {
	vs := VectorRefs{New(1,2,3), New(4,5,6), New(7,8,9)}
	vs2 := VectorRefs{New(9,8,7), New(6,5,4), New(3,2,1)}
	vs.AddAllRefs(vs2)
	if fmt.Sprint(*vs[0], *vs[1], *vs[2]) != "{10 10 10} {10 10 10} {10 10 10}" {
		t.Error(fmt.Sprint(*vs[0], *vs[1], *vs[2]))
	}
}

func TestVecRefsCrossVecRefs(t *testing.T) {
	vs := VectorRefs{New(1,2,3), New(4,5,6), New(7,8,9)}
	vs2 := VectorRefs{New(9,8,7), New(6,5,4), New(3,2,1)}
	vs.CrossAllRefs(vs2)
	if fmt.Sprint(*vs[0], *vs[1], *vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}" {
		t.Error(fmt.Sprint(*vs[0], *vs[1], *vs[2]))
	}
}

func TestVecRefsSelect(t *testing.T) {
	vs := VectorRefs{New(1,2,3), New(4,5,6), New(7,8,9)}
	vs2:= vs.Select(
		func(vr *Vector)bool{
			return vr.y==baseScale(5)
		},
	)
	if vs2[0] != vs[1] {
		t.Error(vs2[0],vs[1])
	}
}

func TestVecRefsStride(t *testing.T) {
	vs := VectorRefs{New(1,2,3), New(4,5,6), New(7,8,9)}
	vs2:=vs.Stride(2)
	if fmt.Sprint(vs2.Dereference()) != "[{1 2 3} {7 8 9}]" {
		t.Error(fmt.Sprint(vs2.Dereference()))
	}
}


func TestVecRefsSplit(t *testing.T) {
	vs := VectorRefs{New(1,2,3), New(4,5,6), New(7,8,9)}
	vs2:= vs.Split(
		func(vr *Vector)uint{
			return uint(baseUnscale(vr.y)-2)
		},
	)
	if fmt.Sprint(vs2) != fmt.Sprint([]VectorRefs{nil,nil,VectorRefs{vs[1]},nil,nil,VectorRefs{vs[2]}}) {
		t.Error(vs2,vs)
	}
	vs3:= vs.Split(
		func(vr *Vector)uint{
			return uint(baseUnscale(vr.z-vr.x)-1)
		},
	)
	if fmt.Sprint(vs3[0]) != fmt.Sprint(vs) {
		t.Error(vs3,vs)
	}
}

func TestVecRefsReginalSlicesInChunks(t *testing.T) {
	Hints.ChunkSizeFixed = true
	defer func(dcs int) {
		Hints.ChunkSizeFixed = false
		Hints.DefaultChunkSize = dcs 
	}(Hints.DefaultChunkSize)
	Hints.DefaultChunkSize = 2

	vrs := VectorRefs{New(1, 2, 3), New(4, 5, 6), New(7, 8, 9), New(10, 11, 12), New(13, 14, 15)}
	var vrs2 []VectorRefs
	
	for vrss:=range vectorRefsInRegionalChunks(vrs,vrs.Middle(),4){
		vrs2=append(vrs2,vrss)
	}
	
	if fmt.Sprint(vrs2) != fmt.Sprintf("[[%p %p %p] [] [] [] [] [] [] [%p %p]]",vrs[0],vrs[1],vrs[2],vrs[3],vrs[4]) {
		t.Error(fmt.Sprint(vrs2),fmt.Sprintf("[[%p %p %p] [] [] [] [] [] [] [%p %p]]",vrs[0],vrs[1],vrs[2],vrs[3],vrs[4]))
	}
	vrs2=vrs2[:0]
}

func BenchmarkVecRefsProduct(b *testing.B) {
	b.StopTimer()
	vrs := make(VectorRefs, 100000)
	for i := range vrs {
		vrs[i] = New(1,2,3)
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
		vrs[i] = New(1,2,3)
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

