package tensor3_test

import "fmt"
import . "github.com/splace/tensor3/32bit"

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

