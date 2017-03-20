package tensor3

import "testing"
import "fmt"

func TestVecRefsNewFromIndexes(t *testing.T) {
	vr:=NewVectorRefsFromIndexes(NewVectors(1, 2, 3, 4, 5, 6,7),1,3,2)
	if fmt.Sprint(*vr[2]) != "{4 5 6}" {
		t.Error(fmt.Sprint(*vr[2]))
	}
}

func TestVecRefsDereference(t *testing.T) {
	vr:=NewVectorRefsFromIndexes(NewVectors(1, 2, 3, 4, 5, 6,7),1,3,2)
	if fmt.Sprint(vr.Dereference()) != "[{1 2 3} {7 0 0} {4 5 6}]" {
		t.Error(fmt.Sprint(vr.Dereference()))
	}
}


func TestVecsFromVectorRefs(t *testing.T) {
	vs:=NewVectors(1, 2, 3, 4, 5, 6,7)
	vr1:=NewVectorRefsFromIndexes(vs,1)
	vr2:=NewVectorRefsFromIndexes(vs,3)
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
	vs:=NewVectors(1, 2, 3, 4, 5, 6,7,8,9,10,11,12,13,14,15,16)
	vr1:=NewVectorRefsFromIndexes(vs,3,5,1)
	vr2:=NewVectorRefsFromIndexes(vs,2,5)
	vr:=NewVectorsFromVectorRefs(vr2,vr1)
	is1:=vr.Indexes(vr1)
	is2:=vr.Indexes(vr2)
	if fmt.Sprint(is1,is2,vr) != "[3 2 4] [1 2] [{4 5 6} {13 14 15} {7 8 9} {1 2 3}]" {
		t.Error(fmt.Sprint(is1,is2,vr))
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




