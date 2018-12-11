package tensor3_test

import "fmt"
import . "github.com/splace/tensor3/scaledint"
import "math"
import "math/rand"

func ExampleTextEncodingVector(){
	v1:=New(1,0,0)
	fmt.Printf("%v",*v1)
	// Output:
	// {1 0 0}
}

func ExampleTextEncodingMatrix(){
	m1:=Matrix{*New(1,0,0),*New(1,0,0),*New(1,0,0)}
	fmt.Printf("%v",m1)
	// Output:
	// [{1 0 0} {1 0 0} {1 0 0}]
}


func ExampleBounds(){
	vs:=NewVectors(1,0,0,0,1,0,0,0,1)
	fmt.Print(vs.Max(),vs.Min())
	// Output:
	// {1 1 1} {0 0 0}
}

func ExampleUnitify(){
	vs:=NewVectors(2,0,0,0,-11,0,0,0,5)
	vs.ForEachNoParameter(
		func(v *Vector){
			v.Divide(Base64(math.Sqrt(float64(v.LengthLength()))))
		},
	)
	fmt.Print(vs)
	// Output:
	// [{1 0 0} {0 -1 0} {0 0 1}]
}


func ExampleForEachVector(){
	vs:=NewVectors(1,0,-3)
	vs.ForEachNoParameter(func(v *Vector){v.Multiply(Base32(2))})
	fmt.Printf("%+v",vs)
	// Output:
	// [{2 0 -6}]
}


func ExampleSmallestSeparation(){
	var rnd = rand.New(rand.NewSource(0))
	vs:=make(Vectors,10000)
	for i := range vs{
		vs[i]=*New(rnd.NormFloat64()*100,rnd.NormFloat64()*100,rnd.NormFloat64()*100)
	}

	separation:=func(v1,v2 Vector) BaseType{
		v1.Subtract(v2)
		return v1.LengthLength()
	}

	//start:=time.Now()
	Parallel=true
	i1,i2:=vs.SearchMin(separation)
	Parallel=false
	//fmt.Printf("%v %v %v %v %v %v %v",il,jl,math.Sqrt(float64(ll)),len(vrs),vrs[il],vrs[jl],time.Since(start))
	fmt.Printf("%v %v %v",i1,i2,math.Sqrt(float64(separation(vs[i1],vs[i2]))))	// Output:
	// 3159 8069 0.56423
}

//TODO smallest area regional spliting


func ExampleBoxArea() {
	boxVertices:=NewVectors(1, 1, 1, 1, -1, 1,-1, 1, 1,-1, -1, 1,1, 1, -1,1, -1, -1,-1, 1, -1,-1, -1, -1)
	boxSurfaceTriStrip := NewVectorRefsFromIndexes(boxVertices,1,2,3,4,7,8,5,6,6,8,8,4,6,2,5,1,7,3)
	areax2:=make(chan float64)
	go func(){
		boxSurfaceTriStrip.ForEachInSlices(3,1,false,
			func(tri VectorRefs) {
				v1:=Vector{}
				v1.Set(*tri[0])
				v1.Subtract(*tri[1])
				v2:=Vector{}
				v2.Set(*tri[0])
				v2.Subtract(*tri[2])
				v1.Cross(v2)
				areax2 <- math.Sqrt(float64(v1.LengthLength()))
			})
		close(areax2)
	}()
	var tAreax2 float64
	for c:=range areax2{
		tAreax2+=c
	}
	fmt.Println(tAreax2/2/math.Sqrt(float64(Base64(1))))
	// Output:
	// 24
}




