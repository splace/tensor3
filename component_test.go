package tensor3_test

import "fmt"
import . "github.com/splace/tensor3"
import "math"
import "math/rand"

//import "time"

func ExampleTextEncodingVector() {
	v1 := New(1, 0, 0)
	fmt.Printf("%+v", *v1)
	// Output:
	// {x:1 y:0 z:0}
}

func ExampleTextEncodingMatrix() {
	m1 := Matrix{*New(1, 0, 0), *New(1, 0, 0), *New(1, 0, 0)}
	fmt.Printf("%+v", m1)
	// Output:
	// [{x:1 y:0 z:0} {x:1 y:0 z:0} {x:1 y:0 z:0}]
}

func ExampleBounds() {
	vs := NewVectors(1, 0, 0, 0, 1, 0, 0, 0, 1)
	fmt.Print(vs.Max(), vs.Min())
	// Output:
	// {1 1 1} {0 0 0}
}

func ExampleUnitify() {
	vs := NewVectors(2, 0, 0, 0, -11, 0, 0, 0, 0.1)
	vs.ForEachNoParameter(
		func(v *Vector) {
			v.Divide(Scalar(math.Sqrt(float64(v.LengthLength()))))
		},
	)
	fmt.Print(vs)
	// Output:
	// [{1 0 0} {0 -1 0} {0 0 1}]
}

func ExampleForEachVector() {
	vs := NewVectors(1, 0, 0)
	vs.ForEachNoParameter(func(v *Vector) { v.Multiply(2) })
	fmt.Printf("%+v", vs)
	// Output:
	// [{x:2 y:0 z:0}]
}

func ExampleSmallestSeparation() {
	var rnd = rand.New(rand.NewSource(0))
	// vs:=make(Vectors,10000)  	// 3159 8069 0.5642342569708744
	// vs:=make(Vectors,100000)  	//  61008 63498 0.14127070120538512
	vs := make(Vectors, 20000) //  10233 18996 0.41657537780113923 61008 63498 0.14127070120538512
	for i := range vs {
		vs[i] = *New(rnd.NormFloat64()*100, rnd.NormFloat64()*100, rnd.NormFloat64()*100)
	}

	// separation squared, still matches with smallest
	separation := func(v1, v2 Vector) Scalar {
		v1.Subtract(v2)
		return v1.LengthLength()
	}

	// find the two points that are closest together
	//start:=time.Now()
	//i1,i2,ll:=vs.SearchMin(separation)
	Parallel = true
	i1, i2 := vs.SearchMin(separation)
	Parallel = false
	//fmt.Printf("%v %v %v %v %v %v %v",il,jl,math.Sqrt(float64(ll)),len(vrs),vrs[il],vrs[jl],time.Since(start))
	fmt.Printf("%v %v %v", i1, i2, math.Sqrt(float64(separation(vs[i1], vs[i2]))))
	// Output:
	// 10233 18996 0.41657537780113923
}

func ExampleSmallestSeparationRegional() {
	var rnd = rand.New(rand.NewSource(0))
	vs := make(Vectors, 10000)
	for i := range vs {
		vs[i] = *New(rnd.NormFloat64()*100, rnd.NormFloat64()*100, rnd.NormFloat64()*100)
	}

	separation := func(v1, v2 Vector) Scalar {
		v1.Subtract(v2)
		return v1.LengthLength()
	}

	//start:=time.Now()

	i1, i2 := vs.SearchMinRegionally(separation)
	//i1,i2,_:=vs.SearchMinRegionallyCentered(Vector{},separation)

	//fmt.Printf("%v %v %v %v %v %v %v",il,jl,math.Sqrt(float64(ll)),len(vrs),vrs[il],vrs[jl],time.Since(start))
	fmt.Printf("%v %v %v", i1, i2, math.Sqrt(float64(separation(vs[i1], vs[i2]))))
	// Output:
	// 3159 8069 0.5642342569708744
}

//TODO smallest separation regional splitting

func ExampleBoxArea() {
	// find surface area of triangle strip set
	boxVertices := NewVectors(1, 1, 1, 1, -1, 1, -1, 1, 1, -1, -1, 1, 1, 1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1)
	boxSurfaceTriStrip := NewVectorRefsFromIndexes(boxVertices, 1, 2, 3, 4, 7, 8, 5, 6, 6, 8, 8, 4, 6, 2, 5, 1, 7, 3)
	areax2 := make(chan float64)
	go func() {
		boxSurfaceTriStrip.ForEachInSlices(3, 1, false,
			func(tri VectorRefs) {
				v1 := Vector{}
				v1.Set(*tri[0])
				v1.Subtract(*tri[1])
				v2 := Vector{}
				v2.Set(*tri[0])
				v2.Subtract(*tri[2])
				v1.Cross(v2)
				areax2 <- math.Sqrt(float64(v1.LengthLength()))
			})
		close(areax2)
	}()
	// add together the areas (doubled) arriving on the channel
	var tAreax2 float64
	for c := range areax2 {
		tAreax2 += c
	}
	fmt.Println(tAreax2 / 2)
	// Output:
	// 24
}
