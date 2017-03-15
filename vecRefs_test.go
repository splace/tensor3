package tensor3

import "testing"
import "fmt"

func TestVecRefsNewFromIndexes(t *testing.T) {
	vr:=NewVectorRefsFromIndexes([]uint{1,3,2},NewVectors(1, 2, 3, 4, 5, 6,7)...)
	if fmt.Sprint(*vr[2]) != "{4 5 6}" {
		t.Error(fmt.Sprint(*vr[2]))
	}
}

func TestVecRefsDereference(t *testing.T) {
	vr:=NewVectorRefsFromIndexes([]uint{1,3,2},NewVectors(1, 2, 3, 4, 5, 6,7)...)
	if fmt.Sprint(vr.Dereference()) != "[{1 2 3} {7 0 0} {4 5 6}]" {
		t.Error(fmt.Sprint(vr.Dereference()))
	}
}


func TestVecsFromVectorRefs(t *testing.T) {
	vs:=NewVectors(1, 2, 3, 4, 5, 6,7)
	vr1:=NewVectorRefsFromIndexes([]uint{1},vs...)
	vr2:=NewVectorRefsFromIndexes([]uint{3},vs...)
	vr:=NewVectorsFromVectorRefs(vr1,vr2)
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
	vs:=NewVectors(1, 2, 3, 4, 5, 6,7)
	vr1:=NewVectorRefsFromIndexes([]uint{3},vs...)
	vr2:=NewVectorRefsFromIndexes([]uint{2},vs...)
	vr:=NewVectorsFromVectorRefs(vr2,vr1)
	if fmt.Sprint(vr) != "[{4 5 6} {7 0 0}]" {
		t.Error(fmt.Sprint(vr))
	}
	is:=vr.Indexes(vr1)
	if fmt.Sprint(is) != "[2]" {
		t.Error(fmt.Sprint(is))
	}

}


func TestVecRefsPrint(t *testing.T) {
	vr := VectorRefs{&Vector{1, 2, 3}}
	if fmt.Sprint(*vr[0]) != "{1 2 3}" {
		t.Error(fmt.Sprint(*vr[0]))
	}
}


func TestVecRefsSum(t *testing.T) {
	vs := VectorRefs{&Vector{7, 8, 9}, &Vector{7, 8, 9}, &Vector{7, 8, 9}}
	if fmt.Sprint(vs.Sum()) != "{21 24 27}" {
		t.Error(vs.Sum())
	}
}




