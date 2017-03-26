package tensor3

import "testing"
import "fmt"

func TestVecRefsNewFromIndexes(t *testing.T) {
	vr:=NewVectorRefsFromIndexes(NewVectors(1, 2, 3, 4, 5, 6,7),1,3,2)
	if fmt.Sprint(*vr[2]) != "{4 5 6}" {
		t.Error(fmt.Sprint(*vr[2]))
	}
}

func TestVecRefsNewFromEmplyIndexes(t *testing.T) {
	vr:=NewVectorRefsFromIndexes(NewVectors(1, 2, 3, 4, 5, 6, 7))
	if fmt.Sprint(*vr[2]) != "{7 0 0}" {
		t.Error(fmt.Sprint(*vr[2]))
	}
}

func TestVecRefsDereferenceNew(t *testing.T) {
	vr:=NewVectorRefsFromIndexes(NewVectors(1, 2, 3, 4, 5, 6,7),1,3,2)
	if fmt.Sprint(vr.Dereference()) != "[{1 2 3} {7 0 0} {4 5 6}]" {
		t.Error(fmt.Sprint(vr.Dereference()))
	}
}

func TestVecRefsDereference(t *testing.T) {
	vr:=NewVectorRefsFromIndexes(NewVectors(1, 2, 3, 4, 5, 6,7),1,3,2)
	v:=make(Vectors,2)
	v.Dereference(vr)
	if fmt.Sprint(v) != "[{1 2 3} {7 0 0}]" {
		t.Error(fmt.Sprint(v))
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
	is1:=vr1.Indexes(vr)
	is2:=vr2.Indexes(vr)
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

func TestVecRefsAddRefs(t *testing.T) {
	vs := VectorRefs{&Vector{1,2, 3}, &Vector{4, 5, 6}, &Vector{7, 8, 9}}
	vs2 := VectorRefs{&Vector{9, 8, 7}, &Vector{6, 5, 4}, &Vector{3, 2, 1}}
	vs.AddAllRefs(vs2)
	if fmt.Sprint(*vs[0],*vs[1],*vs[2]) != "{10 10 10} {10 10 10} {10 10 10}"{
		t.Error(fmt.Sprint(*vs[0],*vs[1],*vs[2]))
	}
}

func TestVecRefsCrossRefs(t *testing.T) {
	vs := VectorRefs{&Vector{1,2, 3}, &Vector{4, 5, 6}, &Vector{7, 8, 9}}
	vs2 := VectorRefs{&Vector{9, 8, 7}, &Vector{6, 5, 4}, &Vector{3, 2, 1}}
	vs.CrossAllRefs(vs2)
	if fmt.Sprint(*vs[0],*vs[1],*vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}"{
		t.Error(fmt.Sprint(*vs[0],*vs[1],*vs[2]))
	}
}


func TestVecsAddVecRefs(t *testing.T) {
	vs := Vectors{Vector{1,2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	vs2 := VectorRefs{&Vector{9, 8, 7}, &Vector{6, 5, 4}, &Vector{3, 2, 1}}
	vs.AddAllRefs(vs2)
	if fmt.Sprint(vs[0],vs[1],vs[2]) != "{10 10 10} {10 10 10} {10 10 10}"{
		t.Error(fmt.Sprint(vs[0],vs[1],vs[2]))
	}
}

func TestVecsCrossVecRefs(t *testing.T) {
	vs := Vectors{Vector{1,2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9}}
	vs2 := VectorRefs{&Vector{9, 8, 7}, &Vector{6, 5, 4}, &Vector{3, 2, 1}}
	vs.CrossAllRefs(vs2)
	if fmt.Sprint(vs[0],vs[1],vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}"{
		t.Error(fmt.Sprint(vs[0],vs[1],vs[2]))
	}
}
func TestVecsRefsAddVecs(t *testing.T) {
	vs := VectorRefs{&Vector{1,2, 3}, &Vector{4, 5, 6}, &Vector{7, 8, 9}}
	vs2 := Vectors{Vector{9, 8, 7}, Vector{6, 5, 4}, Vector{3, 2, 1}}
	vs.AddAll(vs2)
	if fmt.Sprint(*vs[0],*vs[1],*vs[2]) != "{10 10 10} {10 10 10} {10 10 10}"{
		t.Error(fmt.Sprint(*vs[0],*vs[1],*vs[2]))
	}
}

func TestVecRefsCrossVecs(t *testing.T) {
	vs := VectorRefs{&Vector{1,2, 3}, &Vector{4, 5, 6}, &Vector{7, 8, 9}}
	vs2 := Vectors{Vector{9, 8, 7}, Vector{6, 5, 4}, Vector{3, 2, 1}}
	vs.CrossAll(vs2)
	if fmt.Sprint(*vs[0],*vs[1],*vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}"{
		t.Error(fmt.Sprint(*vs[0],*vs[1],*vs[2]))
	}
}

func TestVecsRefsAddVecRefs(t *testing.T) {
	vs := VectorRefs{&Vector{1,2, 3}, &Vector{4, 5, 6}, &Vector{7, 8, 9}}
	vs2 := VectorRefs{&Vector{9, 8, 7}, &Vector{6, 5, 4}, &Vector{3, 2, 1}}
	vs.AddAllRefs(vs2)
	if fmt.Sprint(*vs[0],*vs[1],*vs[2]) != "{10 10 10} {10 10 10} {10 10 10}"{
		t.Error(fmt.Sprint(*vs[0],*vs[1],*vs[2]))
	}
}

func TestVecRefsCrossVecRefs(t *testing.T) {
	vs := VectorRefs{&Vector{1,2, 3}, &Vector{4, 5, 6}, &Vector{7, 8, 9}}
	vs2 := VectorRefs{&Vector{9, 8, 7}, &Vector{6, 5, 4}, &Vector{3, 2, 1}}
	vs.CrossAllRefs(vs2)
	if fmt.Sprint(*vs[0],*vs[1],*vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}"{
		t.Error(fmt.Sprint(*vs[0],*vs[1],*vs[2]))
	}
}



