package tensor3_test

import "fmt"
import . "github.com/splace/tensor3/scaledint"

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


