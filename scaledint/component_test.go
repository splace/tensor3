package tensor3_test

import "fmt"
import . "github.com/splace/tensor3/scaledint"
import "math"

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
	vs.ForEachNoParameter(func(v *Vector){v.Multiply(2)})
	fmt.Printf("%+v",vs)
	// Output:
	// [{x:2 y:0 z:-6}]
}




