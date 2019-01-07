package tensor3

import "testing"
import "fmt"
import "math"

import "math/rand"

var benchSizes = []int{100,10000,40000}  // 40,000,000 needs about 1gb ram

func TestVecsPrint(t *testing.T) {
	v := Vectors{*New(1, 2, 3)}
	if fmt.Sprint(v) != "[{1 2 3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsNew(t *testing.T) {
	v := NewVectors(1, 2, 3, 4, 5, 6, 7)
	if fmt.Sprint(v) != "[{1 2 3} {4 5 6} {7 0 0}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCrossLen1(t *testing.T) {
	v := Vectors{*New(1, 2, 3)}
	v.Cross(*New(4, 5, 6))
	if fmt.Sprint(v) != "[{-3 6 -3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCross(t *testing.T) {
	v := Vectors{*New(1, 2, 3), *New(1, 2, 3)}
	v.Cross(*New(4, 5, 6))
	if fmt.Sprint(v) != "[{-3 6 -3} {-3 6 -3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsProduct(t *testing.T) {
	v := Vectors{*New(1, 2, 3), *New(1, 2, 3)}
	v.Product(Identity)
	if fmt.Sprint(v) != "[{1 2 3} {1 2 3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsSum(t *testing.T) {
	vs := Vectors{*New(7, 8, 9), *New(7, 8, 9), *New(7, 8, 9)}
	if fmt.Sprint(vs.Sum()) != "{21 24 27}" {
		t.Error(vs.Sum())
	}
}

func TestVecsMax(t *testing.T) {
	vs := Vectors{*New(4, -1, 11), *New(7, 2, 9), *New(7, 8, 9)}
	if fmt.Sprint(vs.Max()) != "{7 8 11}" {
		t.Error(vs.Max())
	}
}

func TestVecsMaxOne(t *testing.T) {
	vs := Vectors{*New(1, -1, 1)}
	if fmt.Sprint(vs.Max()) != "{1 -1 1}" {
		t.Error(vs.Max())
	}
}

func TestVecsMaxNone(t *testing.T) {
	vs := Vectors{}
	defer func() {
		r := recover()
		if r == nil {
			t.Error("Expected error not present.")
		}
	}()
	_ = vs.Max()
}

func TestVecsMin(t *testing.T) {
	vs := Vectors{*New(4, -1, 11), *New(7, 2, 9), *New(7, 8, 9)}
	if fmt.Sprint(vs.Min()) != "{4 -1 9}" {
		t.Error(vs.Min())
	}
}

func TestVecsMiddle(t *testing.T) {
	vs := Vectors{*New(4, -1, 11), *New(7, 2, 9), *New(7, 8, 9)}
	if fmt.Sprint(vs.Middle()) != "{5.5 3.5 10}" {
		t.Error(vs.Middle())
	}
}

func TestVecsInterpolate(t *testing.T) {
	vs := Vectors{*New(7, 8, 9), *New(7, 8, 9), *New(7, 8, 9)}
	vs.Interpolate(Vector{-2, 1, -1}, 0.5)
	if fmt.Sprint(vs) != "[{2.5 4.5 4} {2.5 4.5 4} {2.5 4.5 4}]" {
		t.Error(vs)
	}
}

func TestVecsExtrapolate(t *testing.T) {
	vs := Vectors{*New(7, 8, 9), *New(7, 8, 9), *New(7, 8, 9)}
	vs.Interpolate(*New(-2, 1, -1), 2)
	if fmt.Sprint(vs) != "[{16 15 19} {16 15 19} {16 15 19}]" {
		t.Error(vs)
	}
}

func TestVecsProductT(t *testing.T) {
	v := Vectors{*New(1, 2, 3), *New(1, 2, 3)}
	v.ProductT(Identity)
	if fmt.Sprint(v) != "[{1 2 3} {1 2 3}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCrossChunked1(t *testing.T) {
	Parallel = true
	defer func(d int) {
		Parallel = false
		Hints.DefaultChunkSize = d
	}(Hints.DefaultChunkSize)
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 1
	v := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9), *New(10, 11, 12)}
	//	v=append(v,Vector{1, 2, 3})
	v.Add(*New(1, 1, 1))
	if fmt.Sprint(v) != "[{2 3 4} {5 6 7} {8 9 10} {11 12 13}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCrossChunked2(t *testing.T) {
	Parallel = true
	defer func(d int) {
		Parallel = false
		Hints.DefaultChunkSize = d
	}(Hints.DefaultChunkSize)
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 2
	v := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9), *New(10, 11, 12)}
	//	v=append(v,Vector{1, 2, 3})
	v.Add(*New(1, 1, 1))
	if fmt.Sprint(v) != "[{2 3 4} {5 6 7} {8 9 10} {11 12 13}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsCrossChunked3(t *testing.T) {
	Parallel = true
	defer func(d int) {
		Parallel = false
		Hints.DefaultChunkSize = d
	}(Hints.DefaultChunkSize)
	Hints.ChunkSizeFixed = true
	Hints.DefaultChunkSize = 2
	v := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9), *New(10, 11, 12), *New(13, 14, 15), *New(16, 17, 18)}
	//	v=append(v,Vector{1, 2, 3})
	v.Add(*New(1, 1, 1))
	if fmt.Sprint(v) != "[{2 3 4} {5 6 7} {8 9 10} {11 12 13} {14 15 16} {17 18 19}]" {
		t.Error(fmt.Sprint(v))
	}
}

func TestVecsAddVecs(t *testing.T) {
	vs := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9)}
	vs2 := Vectors{*New(9, 8, 7), *New(6, 5, 4), *New(3, 2, 1)}
	vs.AddAll(vs2)
	if fmt.Sprint(vs[0], vs[1], vs[2]) != "{10 10 10} {10 10 10} {10 10 10}" {
		t.Error(fmt.Sprint(vs[0], vs[1], vs[2]))
	}
}

func TestVecsCrossVecs(t *testing.T) {
	vs := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9)}
	vs2 := Vectors{*New(9, 8, 7), *New(6, 5, 4), *New(3, 2, 1)}
	vs.CrossAll(vs2)
	if fmt.Sprint(vs[0], vs[1], vs[2]) != "{-10 20 -10} {-10 20 -10} {-10 20 -10}" {
		t.Error(fmt.Sprint(vs[0], vs[1], vs[2]))
	}
}

func TestVecsSlicesInChunks(t *testing.T) {
	Hints.ChunkSizeFixed = true
	defer func(dcs int) {
		Hints.ChunkSizeFixed = false
		Hints.DefaultChunkSize = dcs
	}(Hints.DefaultChunkSize)
	Hints.DefaultChunkSize = 2

	vs := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9), *New(10, 11, 12), *New(13, 14, 15)}
	var vs2 [][]Vectors

	for vss := range vectorSlicesInChunks(vs, 10, 1, 1, false) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3}] [{4 5 6}] [{7 8 9}] [{10 11 12}] [{13 14 15}]]]" {
		t.Error(fmt.Println(vs2))
	}
	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 10, 2, 1, false) {
		vs2 = append(vs2, vss)
	}

	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6}] [{4 5 6} {7 8 9}] [{7 8 9} {10 11 12}] [{10 11 12} {13 14 15}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 10, 3, 1, false) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6} {7 8 9}] [{4 5 6} {7 8 9} {10 11 12}] [{7 8 9} {10 11 12} {13 14 15}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 1, 1, 1, false) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3}]] [[{4 5 6}]] [[{7 8 9}]] [[{10 11 12}]] [[{13 14 15}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 2, 2, 1, false) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6}] [{4 5 6} {7 8 9}]] [[{7 8 9} {10 11 12}] [{10 11 12} {13 14 15}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 4, 2, 1, false) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6}] [{4 5 6} {7 8 9}] [{7 8 9} {10 11 12}] [{10 11 12} {13 14 15}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 3, 1, 1, false) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3}] [{4 5 6}] [{7 8 9}]] [[{10 11 12}] [{13 14 15}]]]" {
		t.Error(fmt.Println(vs2))
	}

}

func TestVecsSlicesStridingInChunks(t *testing.T) {
	Hints.ChunkSizeFixed = true
	defer func(dcs int) {
		Hints.ChunkSizeFixed = false
		Hints.DefaultChunkSize = dcs
	}(Hints.DefaultChunkSize)
	Hints.DefaultChunkSize = 2

	vs := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9), *New(10, 11, 12), *New(13, 14, 15), *New(16, 17, 18), *New(19, 20, 21), *New(22, 23, 24)}

	var vs2 [][]Vectors

	for vss := range vectorSlicesInChunks(vs, 10, 1, 2, false) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3}] [{7 8 9}] [{13 14 15}] [{19 20 21}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 10, 2, 3, false) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6}] [{10 11 12} {13 14 15}] [{19 20 21} {22 23 24}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 4, 1, 2, false) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3}] [{7 8 9}]] [[{13 14 15}] [{19 20 21}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 4, 2, 2, false) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6}] [{7 8 9} {10 11 12}]] [[{13 14 15} {16 17 18}] [{19 20 21} {22 23 24}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 4, 2, 1, false) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6}] [{4 5 6} {7 8 9}] [{7 8 9} {10 11 12}] [{10 11 12} {13 14 15}]] [[{13 14 15} {16 17 18}] [{16 17 18} {19 20 21}] [{19 20 21} {22 23 24}]]]" {
		t.Error(fmt.Println(vs2))
	}

}

func TestVecsSlicesStridingAndWrappingInChunks(t *testing.T) {
	Hints.ChunkSizeFixed = true
	defer func(dcs int) {
		Hints.ChunkSizeFixed = false
		Hints.DefaultChunkSize = dcs
	}(Hints.DefaultChunkSize)
	Hints.DefaultChunkSize = 2

	vs := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9), *New(10, 11, 12), *New(13, 14, 15), *New(16, 17, 18), *New(19, 20, 21), *New(22, 23, 24)}

	var vs2 [][]Vectors

	for vss := range vectorSlicesInChunks(vs, 10, 1, 2, true) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3}] [{7 8 9}] [{13 14 15}] [{19 20 21}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 10, 2, 3, true) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6}] [{10 11 12} {13 14 15}] [{19 20 21} {22 23 24}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 4, 1, 2, true) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3}] [{7 8 9}]] [[{13 14 15}] [{19 20 21}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 4, 2, 2, true) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6}] [{7 8 9} {10 11 12}]] [[{13 14 15} {16 17 18}] [{19 20 21} {22 23 24}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 4, 2, 3, true) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6}] [{10 11 12} {13 14 15}]] [[{19 20 21} {22 23 24}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 10, 2, 3, true) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6}] [{10 11 12} {13 14 15}] [{19 20 21} {22 23 24}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 4, 3, 2, true) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6} {7 8 9}] [{7 8 9} {10 11 12} {13 14 15}]] [[{13 14 15} {16 17 18} {19 20 21}] [{19 20 21} {22 23 24} {1 2 3}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 10, 3, 1, true) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6} {7 8 9}] [{4 5 6} {7 8 9} {10 11 12}] [{7 8 9} {10 11 12} {13 14 15}] [{10 11 12} {13 14 15} {16 17 18}] [{13 14 15} {16 17 18} {19 20 21}] [{16 17 18} {19 20 21} {22 23 24}] [{19 20 21} {22 23 24} {1 2 3}] [{22 23 24} {1 2 3} {4 5 6}]]]" {
		t.Error(fmt.Println(vs2))
	}

	vs2 = vs2[:0]
	for vss := range vectorSlicesInChunks(vs, 6, 3, 1, true) {
		vs2 = append(vs2, vss)
	}
	if fmt.Sprint(vs2) != "[[[{1 2 3} {4 5 6} {7 8 9}] [{4 5 6} {7 8 9} {10 11 12}] [{7 8 9} {10 11 12} {13 14 15}] [{10 11 12} {13 14 15} {16 17 18}] [{13 14 15} {16 17 18} {19 20 21}] [{16 17 18} {19 20 21} {22 23 24}] [{19 20 21} {22 23 24} {1 2 3}] [{22 23 24} {1 2 3} {4 5 6}]]]" {
		t.Error(fmt.Println(vs2))
	}

}







func TestVecsvectorsFindMin(t *testing.T) {
	vs:=Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9), *New(10, 11, 12), *New(13, 14, 15)}
	i := vectorsFindMin(
		vs,
		func(v Vector) Scalar { return -v.x },
	)
	if i != 4 {
		t.Error()
	}

}

func TestVecsvectorsFindMinChunked(t *testing.T) {
	vs:=Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(7, 8, 9), *New(10, 11, 12), *New(13, 14, 15)}
	i := vectorsFindMinChunked(
		vs,
		func(v Vector) Scalar { return -v.x },
	)
	if i != 4 {
		t.Error()
	}

}

func TestVecsSearchMinRegional(t *testing.T) {
	vs := Vectors{*New(1, 1, 1), *New(1, 1, -1), *New(1, -1, -1), *New(1, -1, 1), *New(-1, 1, 1), *New(-1, 1, -1), *New(-1, -1, -1), *New(-1, -1, 1)}

	separation := func(v1, v2 Vector) Scalar {
		return v1.DistDist(v2)
	}

	iv, jv := vs.SearchMinRegionally(separation)

	// algorithm used returned first match (not necesserily if parallel.)
	if iv != &vs[0] || jv != &vs[1] {
		t.Error(iv, jv)
	}else{
		if math.Sqrt(float64(separation(*iv, *jv))) != 2 {
			t.Error(iv, jv, math.Sqrt(float64(separation(*iv, *jv))))
		}
	}
}
	

func TestVecsSearchMindX(t *testing.T) {
	vs := Vectors{*New(1, 2, 3), *New(4, 5, 6), *New(9, 8, 9), *New(10, 11, 12), *New(13, 14, 15)}
	i, j := vs.SearchMin(
		func(v1, v2 Vector) Scalar {
			return v2.x - v1.x
		},
	)
	if i != 2 || j != 3 {
		t.Error(i, j)
	}
	i, j = vs.SearchMin(
		func(v1, v2 Vector) Scalar {
			return v1.x - v2.x
		},
	)
	if i != 0 || j != 4 {
		t.Error(i, j)
	}
}

func TestVecsSearchMin(t *testing.T) {
	var rnd = rand.New(rand.NewSource(0))
	vs := make(Vectors, 10000)
	for i := range vs {
		vs[i] = *New(rnd.NormFloat64()*100, rnd.NormFloat64()*100, rnd.NormFloat64()*100)
	}

	separation := func(v1, v2 Vector) Scalar {
		v1.Subtract(v2)
		return v1.LengthLength()
	}

	i, j := vs.SearchMin(separation)

	if i != 3159 || j != 8069 {
		t.Error(i, j)
	}else{
		if math.Sqrt(float64(separation(vs[i], vs[j]))) != 0.5642342569708744 {
			t.Error(i, j, math.Sqrt(float64(separation(vs[i], vs[j]))))
		}
	}
	
	iv, jv := vs.SearchMinRegionally(separation)

	if iv != &vs[3159] || jv != &vs[8069] {
		t.Error(iv, jv)
	}else{
		if math.Sqrt(float64(separation(*iv, *jv))) != 0.5642342569708744 {
			t.Error(iv, jv, math.Sqrt(float64(separation(*iv, *jv))))
		}
	}
}



func TestVecsForEachInSlices(t *testing.T) {
	vs := Vectors{*New(0, 0, 0), *New(1, 0, 0), *New(1, 1, 0)}

	var c float64
	vs.ForEachInSlices(1, 1, false,
		func(vss Vectors) {
			vss[0] = *New(c, c+1, c+2)
			c += 3
		})
	if fmt.Sprint(vs) != "[{0 1 2} {3 4 5} {6 7 8}]" {
		t.Error(fmt.Println(vs))
	}
	vs.ForEachInSlices(2, 1, false,
		func(vss Vectors) {
			vss[0] = *New(c, c+1, c+2)
			c += 3
		})

	// doesn't attemp to update wrapped
	if fmt.Sprint(vs) != "[{9 10 11} {12 13 14} {6 7 8}]" {
		t.Error(fmt.Println(vs))
	}

	vs.ForEachInSlices(2, 1, true,
		func(vss Vectors) {
			vss[0] = *New(c, c+1, c+2)
			c += 3
		})

	if fmt.Sprint(vs) != "[{15 16 17} {18 19 20} {21 22 23}]" {
		t.Error(fmt.Println(vs))
	}

	vs.ForEachInSlices(1, 2, true,
		func(vss Vectors) {
			vss[0] = *New(c, c+1, c+2)
			c += 3
		})

	if fmt.Sprint(vs) != "[{24 25 26} {18 19 20} {27 28 29}]" {
		t.Error(fmt.Println(vs))
	}

	vs.ForEachInSlices(2, 2, true,
		func(vss Vectors) {
			vss[0] = *New(c, c+1, c+2)
			vss[1] = *New(c+3, c+4, c+5)
			c += 6
		})

	if fmt.Sprint(vs) != "[{39 40 41} {33 34 35} {36 37 38}]" {
		t.Error(fmt.Println(vs))
	}

	vs.ForEachInSlices(2, 3, false,
		func(vss Vectors) {
			vss[0] = *New(c, c+1, c+2)
			vss[1] = *New(c+3, c+4, c+5)
			c += 6
		})

	if fmt.Sprint(vs) != "[{42 43 44} {45 46 47} {36 37 38}]" {
		t.Error(fmt.Println(vs))
	}

}

//TODO parallel strip area

func TestVecsTriangleStripArea(t *testing.T) {
	vs := Vectors{*New(0, 0, 0), *New(1, 0, 0), *New(0, 2, 0), *New(1, 2, 0)}
	areax2 := make(chan float64)
	go func() {
		vs.ForEachInSlices(3, 1, false,
			func(tri Vectors) {
				v1 := Vector{}
				v1.Set(tri[0])
				v1.Subtract(tri[1])
				v2 := Vector{}
				v2.Set(tri[0])
				v2.Subtract(tri[2])
				v1.Cross(v2)
				areax2 <- math.Sqrt(float64(v1.LengthLength()))
			})
		close(areax2)
	}()
	var tAreax2 float64
	for c := range areax2 {
		tAreax2 += c
	}
	if tAreax2/math.Sqrt(scale) != 4 {
		t.Error(tAreax2)
	}
}

func BenchmarkVecsSum(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 100000)
	for i := range vs {
		vs[i] = *New(1, 2, 3)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Sum()
	}
}

func BenchmarkVecsSumParallel(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 100000)
	for i := range vs {
		vs[i] = *New(1, 2, 3)
	}
	Parallel = true
	Hints.ChunkSizeFixed = true
	defer func() {
		Parallel = false
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Sum()
	}
}

func BenchmarkVecsCross(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 100000)
	for i := range vs {
		vs[i] = *New(1, 2, 3)
	}
	v := Vector{9, 8, 7}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Cross(v)
	}

}

func BenchmarkVecsCrossParallel(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, 10000)
	for i := range vs {
		vs[i] = *New(1, 2, 3)
	}
	v := *New(9, 8, 7)
	Parallel = true
	Hints.ChunkSizeFixed = true
	defer func() {
		Parallel = false
	}()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		vs.Cross(v)
	}

}

func BenchmarkVecsProduct(b *testing.B) {
	b.StopTimer()
	vs := make(Vectors, benchSizes[len(benchSizes)-1]) 
	for i := range vs {
		vs[i] = *New(1, 2, 3)
	}
	m := Matrix{}
	b.StartTimer()
	for _,e:=range(benchSizes){
        b.Run(fmt.Sprint(e), func(b *testing.B) {
        	vss:=vs[0:e]
		 	for l := 0; l < b.N; l++ {
				vss.Product(m)
				}
			 })
        }
   	b.StopTimer()
	Parallel = true
	Hints.ChunkSizeFixed = true
	defer func() {
		Parallel = false
	}()
	b.StartTimer()
	for _,e:=range(benchSizes){
        b.Run(fmt.Sprint(e)+" Parallel", func(b *testing.B) {
        	vss:=vs[0:e]
		 	for l := 0; l < b.N; l++ {
				vss.Product(m)
				}
			 })
        }

}


func TestVecsFindNearest(t *testing.T) {
	var rnd = rand.New(rand.NewSource(0))
	vs := make(Vectors, 1000)
	for i := range vs {
		vs[i] = *New(rnd.NormFloat64()*100, rnd.NormFloat64()*100, rnd.NormFloat64()*100)
	}

	// distance squared, still finds nearest
	distance := func(v Vector) Scalar {
		return v.LengthLength()
	}

	v:=vs.FindMin(distance)
	if v != 587 ||  math.Sqrt(float64(distance(vs[v])))!=7.261516613941612{
		t.Error(v, math.Sqrt(float64(distance(vs[v]))))
	}

}


func BenchmarkVecsFindNearest(b *testing.B) {
	b.StopTimer()
	var rnd = rand.New(rand.NewSource(0))
	vs := make(Vectors, benchSizes[len(benchSizes)-1]) 
	for i := range vs {
		vs[i] = *New(rnd.NormFloat64()*100, rnd.NormFloat64()*100, rnd.NormFloat64()*100)
	}

	// distance squared, still finds nearest
	distance := func(v Vector) Scalar {
		return v.LengthLength()
	}

	b.StartTimer()
	for _,e:=range(benchSizes){
        b.Run(fmt.Sprint(e), func(b *testing.B) {
        	vss:=vs[0:e]
		 	for l := 0; l < b.N; l++ {
				vss.FindMin(distance)
				}
			 })
        }

	b.StopTimer()
	Parallel = true
	Hints.ChunkSizeFixed = true
	defer func() {
		Parallel = false
	}()
	b.StartTimer()
	for _,e:=range(benchSizes){
        b.Run(fmt.Sprint(e)+" Parallel", func(b *testing.B) {
        	vss:=vs[0:e]
		 	for l := 0; l < b.N; l++ {
				vss.FindMin(distance)
				}
			 })
        }
}


func BenchmarkVecsSearchNearest(b *testing.B) {
	b.StopTimer()
	var rnd = rand.New(rand.NewSource(0))
	vs := make(Vectors, benchSizes[len(benchSizes)-1]) 
	for i := range vs {
		vs[i] = *New(rnd.NormFloat64()*100, rnd.NormFloat64()*100, rnd.NormFloat64()*100)
	}

	separation := func(v1, v2 Vector) Scalar {
		v1.Subtract(v2)
		return v1.LengthLength()
	}

	b.StartTimer()
	for _,e:=range(benchSizes){
        b.Run(fmt.Sprint(e), func(b *testing.B) {
        	vss:=vs[0:e]
		 	for l := 0; l < b.N; l++ {
		 		vss.SearchMin(separation)
				}
			 })
        }

	b.StopTimer()
	Parallel = true
	Hints.ChunkSizeFixed = true
	defer func() {
		Parallel = false
	}()
	b.StartTimer()
	for _,e:=range(benchSizes){
        b.Run(fmt.Sprint(e)+" Parallel", func(b *testing.B) {
        	vss:=vs[0:e]
		 	for l := 0; l < b.N; l++ {
		 		vss.SearchMin(separation)
				}
			 })
        }

	b.StartTimer()
	for _,e:=range(benchSizes){
        b.Run(fmt.Sprint(e)+" Regional", func(b *testing.B) {
        	vss:=vs[0:e]
		 	for l := 0; l < b.N; l++ {
		 		vss.SearchMinRegionally(separation)
				}
			 })
        }

	b.StopTimer()
	Parallel = true
	b.StartTimer()
	for _,e:=range(benchSizes){
        b.Run(fmt.Sprint(e)+" Parallel Regional", func(b *testing.B) {
        	vss:=vs[0:e]
		 	for l := 0; l < b.N; l++ {
		 		vss.SearchMinRegionally(separation)
				}
			 })
        }
}





/* benchmark: "VecsSearchNearest" hal3 Mon 7 Jan 21:52:42 GMT 2019 go version go1.11.4 linux/amd64
goos: linux
goarch: amd64
BenchmarkVecsSearchNearest/100           	   10000	    100260 ns/op
BenchmarkVecsSearchNearest/100-2         	   20000	    100239 ns/op
BenchmarkVecsSearchNearest/10000         	       2	 833960419 ns/op
BenchmarkVecsSearchNearest/10000-2       	       2	 841831106 ns/op
BenchmarkVecsSearchNearest/40000         	       1	13459721197 ns/op
BenchmarkVecsSearchNearest/40000-2       	       1	13586018301 ns/op
BenchmarkVecsSearchNearest/100_Parallel           	   20000	     99814 ns/op
BenchmarkVecsSearchNearest/100_Parallel-2         	   10000	    108897 ns/op
BenchmarkVecsSearchNearest/10000_Parallel         	       2	 847642738 ns/op
BenchmarkVecsSearchNearest/10000_Parallel-2       	       2	 843494601 ns/op
BenchmarkVecsSearchNearest/40000_Parallel         	       1	11921844000 ns/op
BenchmarkVecsSearchNearest/40000_Parallel-2       	       1	14315922120 ns/op
BenchmarkVecsSearchNearest/100_Regional           	   20000	     70598 ns/op
BenchmarkVecsSearchNearest/100_Regional-2         	   20000	     95587 ns/op
BenchmarkVecsSearchNearest/10000_Regional         	      10	 152239460 ns/op
BenchmarkVecsSearchNearest/10000_Regional-2       	      10	 150845765 ns/op
BenchmarkVecsSearchNearest/40000_Regional         	       1	8558573110 ns/op
BenchmarkVecsSearchNearest/40000_Regional-2       	       1	9871147567 ns/op
BenchmarkVecsSearchNearest/100_Parallel_Regional           	   20000	     71665 ns/op
BenchmarkVecsSearchNearest/100_Parallel_Regional-2         	   20000	     95430 ns/op
BenchmarkVecsSearchNearest/10000_Parallel_Regional         	      10	 153071747 ns/op
BenchmarkVecsSearchNearest/10000_Parallel_Regional-2       	      10	 151463460 ns/op
BenchmarkVecsSearchNearest/40000_Parallel_Regional         	       1	8796931835 ns/op
BenchmarkVecsSearchNearest/40000_Parallel_Regional-2       	       1	9805282436 ns/op
PASS
ok  	_/run/media/simon/6a5530c2-1442-4e9b-b35f-3db0c9a6984c/home/simon/Dropbox/github/working/tensor3	125.410s
Mon 7 Jan 21:54:48 GMT 2019
*/
/* benchmark: "VecsSearchNearest" hal3 Mon 7 Jan 21:57:43 GMT 2019 go version go1.11.4 linux/amd64
FAIL	_/run/media/simon/6a5530c2-1442-4e9b-b35f-3db0c9a6984c/home/simon/Dropbox/github/working/tensor3 [build failed]
Mon 7 Jan 21:57:44 GMT 2019
*/
/* benchmark: "VecsSearchNearest" hal3 Mon 7 Jan 21:58:27 GMT 2019 go version go1.11.4 linux/amd64
goos: linux
goarch: amd64
BenchmarkVecsSearchNearest/100           	   20000	     99161 ns/op
BenchmarkVecsSearchNearest/100-2         	   10000	    104249 ns/op
BenchmarkVecsSearchNearest/10000         	       2	 833303773 ns/op
BenchmarkVecsSearchNearest/10000-2       	       2	 836447981 ns/op
BenchmarkVecsSearchNearest/40000         	       1	13115215830 ns/op
BenchmarkVecsSearchNearest/40000-2       	       1	13378893114 ns/op
BenchmarkVecsSearchNearest/100_Parallel           	   10000	    101130 ns/op
BenchmarkVecsSearchNearest/100_Parallel-2         	   10000	    102661 ns/op
BenchmarkVecsSearchNearest/10000_Parallel         	       2	 828924475 ns/op
BenchmarkVecsSearchNearest/10000_Parallel-2       	       2	 831199573 ns/op
BenchmarkVecsSearchNearest/40000_Parallel         	       1	11693743469 ns/op
BenchmarkVecsSearchNearest/40000_Parallel-2       	       1	14002427341 ns/op
BenchmarkVecsSearchNearest/100_Regional           	   20000	     70281 ns/op
BenchmarkVecsSearchNearest/100_Regional-2         	   20000	    101509 ns/op
BenchmarkVecsSearchNearest/10000_Regional         	      10	 149422691 ns/op
BenchmarkVecsSearchNearest/10000_Regional-2       	      10	 151031023 ns/op
BenchmarkVecsSearchNearest/40000_Regional         	       1	8317373105 ns/op
BenchmarkVecsSearchNearest/40000_Regional-2       	       1	9627455505 ns/op
BenchmarkVecsSearchNearest/100_Parallel_Regional           	   20000	     70690 ns/op
BenchmarkVecsSearchNearest/100_Parallel_Regional-2         	   20000	     96563 ns/op
BenchmarkVecsSearchNearest/10000_Parallel_Regional         	      10	 152525535 ns/op
BenchmarkVecsSearchNearest/10000_Parallel_Regional-2       	      10	 150662934 ns/op
BenchmarkVecsSearchNearest/40000_Parallel_Regional         	       1	8524565140 ns/op
BenchmarkVecsSearchNearest/40000_Parallel_Regional-2       	       1	9556594482 ns/op
PASS
ok  	_/run/media/simon/6a5530c2-1442-4e9b-b35f-3db0c9a6984c/home/simon/Dropbox/github/working/tensor3	120.955s
Mon 7 Jan 22:00:29 GMT 2019
*/
/* benchmark: "VecsSearchNearest" hal3 Mon 7 Jan 22:05:40 GMT 2019 go version go1.11.4 linux/amd64
goos: linux
goarch: amd64
BenchmarkVecsSearchNearest/100           	   20000	     98965 ns/op
BenchmarkVecsSearchNearest/100-2         	   10000	    103102 ns/op
BenchmarkVecsSearchNearest/10000         	       2	 833915983 ns/op
BenchmarkVecsSearchNearest/10000-2       	       2	 847039550 ns/op
BenchmarkVecsSearchNearest/40000         	       1	13261212108 ns/op
BenchmarkVecsSearchNearest/40000-2       	       1	13485564474 ns/op
BenchmarkVecsSearchNearest/100_Parallel           	   20000	     98699 ns/op
BenchmarkVecsSearchNearest/100_Parallel-2         	   10000	    100937 ns/op
BenchmarkVecsSearchNearest/10000_Parallel         	       2	 841127772 ns/op
BenchmarkVecsSearchNearest/10000_Parallel-2       	       2	 834027031 ns/op
BenchmarkVecsSearchNearest/40000_Parallel         	       1	11346995150 ns/op
BenchmarkVecsSearchNearest/40000_Parallel-2       	       1	14294683039 ns/op
BenchmarkVecsSearchNearest/100_Regional           	   20000	     70342 ns/op
BenchmarkVecsSearchNearest/100_Regional-2         	   20000	     96772 ns/op
BenchmarkVecsSearchNearest/10000_Regional         	      10	 149138148 ns/op
BenchmarkVecsSearchNearest/10000_Regional-2       	      10	 150826212 ns/op
BenchmarkVecsSearchNearest/40000_Regional         	       1	8301433197 ns/op
BenchmarkVecsSearchNearest/40000_Regional-2       	       1	9493703227 ns/op
BenchmarkVecsSearchNearest/100_Parallel_Regional           	   20000	     71095 ns/op
BenchmarkVecsSearchNearest/100_Parallel_Regional-2         	   20000	     99238 ns/op
BenchmarkVecsSearchNearest/10000_Parallel_Regional         	      10	 152296017 ns/op
BenchmarkVecsSearchNearest/10000_Parallel_Regional-2       	      10	 150084507 ns/op
BenchmarkVecsSearchNearest/40000_Parallel_Regional         	       1	8286578345 ns/op
BenchmarkVecsSearchNearest/40000_Parallel_Regional-2       	       1	9492659896 ns/op
PASS
ok  	_/run/media/simon/6a5530c2-1442-4e9b-b35f-3db0c9a6984c/home/simon/Dropbox/github/working/tensor3	122.972s
Mon 7 Jan 22:07:44 GMT 2019
*/
/* benchmark: "VecsSearchNearest" hal3 Mon 7 Jan 22:39:08 GMT 2019 go version go1.11.4 linux/amd64
goos: linux
goarch: amd64
BenchmarkVecsSearchNearest/100           	   20000	     99542 ns/op
BenchmarkVecsSearchNearest/100-2         	   10000	    103933 ns/op
BenchmarkVecsSearchNearest/10000         	       2	 997626307 ns/op
BenchmarkVecsSearchNearest/10000-2       	       2	1220234282 ns/op
BenchmarkVecsSearchNearest/40000         	       1	15284346302 ns/op
BenchmarkVecsSearchNearest/40000-2       	       1	14203999034 ns/op
BenchmarkVecsSearchNearest/100_Parallel           	   10000	    102911 ns/op
BenchmarkVecsSearchNearest/100_Parallel-2         	   10000	    103082 ns/op
BenchmarkVecsSearchNearest/10000_Parallel         	       2	1027417591 ns/op
BenchmarkVecsSearchNearest/10000_Parallel-2       	       2	 984562080 ns/op
BenchmarkVecsSearchNearest/40000_Parallel         	       1	12565438692 ns/op
BenchmarkVecsSearchNearest/40000_Parallel-2       	       1	21101689385 ns/op
BenchmarkVecsSearchNearest/100_Regional           	   20000	     99760 ns/op
BenchmarkVecsSearchNearest/100_Regional-2         	   10000	    125282 ns/op
BenchmarkVecsSearchNearest/10000_Regional         	      10	 201892372 ns/op
BenchmarkVecsSearchNearest/10000_Regional-2       	      10	 168905771 ns/op
BenchmarkVecsSearchNearest/40000_Regional         	       1	11564739960 ns/op
BenchmarkVecsSearchNearest/40000_Regional-2       	       1	10524300942 ns/op
BenchmarkVecsSearchNearest/100_Parallel_Regional           	   20000	     70204 ns/op
BenchmarkVecsSearchNearest/100_Parallel_Regional-2         	   20000	     94873 ns/op
BenchmarkVecsSearchNearest/10000_Parallel_Regional         	      10	 153832795 ns/op
BenchmarkVecsSearchNearest/10000_Parallel_Regional-2       	      10	 151876860 ns/op
BenchmarkVecsSearchNearest/40000_Parallel_Regional         	       1	8508940537 ns/op
BenchmarkVecsSearchNearest/40000_Parallel_Regional-2       	       1	9777834671 ns/op
PASS
ok  	_/run/media/simon/6a5530c2-1442-4e9b-b35f-3db0c9a6984c/home/simon/Dropbox/github/working/tensor3	140.511s
Mon 7 Jan 22:41:32 GMT 2019
*/

