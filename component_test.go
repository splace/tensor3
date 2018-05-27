package tensor3_test

import "fmt"
import . "github.com/splace/tensor3"
import "math"
//import "math/rand"
//import "time"

func ExampleTextEncodingVector(){
	v1:=New(1,0,0)
	fmt.Printf("%+v",*v1)
	// Output:
	// {x:1 y:0 z:0}
}

func ExampleTextEncodingMatrix(){
	m1:=Matrix{*New(1,0,0),*New(1,0,0),*New(1,0,0)}
	fmt.Printf("%+v",m1)
	// Output:
	// [{x:1 y:0 z:0} {x:1 y:0 z:0} {x:1 y:0 z:0}]
}

func ExampleBounds(){
	vs:=NewVectors(1,0,0,0,1,0,0,0,1)
	fmt.Print(vs.Max(),vs.Min())
	// Output:
	// {1 1 1} {0 0 0}
}

func ExampleUnitify(){
	vs:=NewVectors(2,0,0,0,-11,0,0,0,0.1)
	vs.ForEachNoParameter(
		func(v *Vector){
			v.Divide(BaseType(math.Sqrt(float64(v.LengthLength()))))
		},
	)
	fmt.Print(vs)
	// Output:
	// [{1 0 0} {0 -1 0} {0 0 1}]
}


func ExampleForEachVector(){
	vs:=NewVectors(1,0,0)
	vs.ForEachNoParameter(func(v *Vector){v.Multiply(2)})
	fmt.Printf("%+v",vs)
	// Output:
	// [{x:2 y:0 z:0}]
}



//func ExampleSmallestSeparation(){
//	vrs:=make(Vectors,100000)
//	for i := range vrs{
//		vrs[i]=*New(rand.NormFloat64()*100,rand.NormFloat64()*100,rand.NormFloat64()*100)
//	}
//
//	separation:=func(v1,v2 Vector) BaseType{
//		v1.Subtract(v2)
//		return v1.LengthLength()
//	}
//
//	start:=time.Now()
//	il,jl,ll:=vrs.SearchMin(separation)
//	fmt.Printf("%v %v %v %v %v %v %v",il,jl,math.Sqrt(ll),len(vrs),vrs[il],vrs[jl],time.Since(start))
//	// Output:
//	// [{x:2 y:0 z:0}]
//}



