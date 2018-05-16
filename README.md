/*

Vector, (3-component), and Matrix, (3-Vector) Types.

Useful methods on same, operating in-place ( not returning reference, meaning no single line chained functions.)

Arrays of both, (called Vectors and Matrices), with their useful methods.

single and parallel threaded, can be selected with a global var.

VectorRefs; array of Vector pointers, with methods to convert to/from Vectors.

64bit float, 32bit float, or fixed precision (int scaled for 3dp) component value types (separate packages).

doesn't use "math" package, left to importers, if necessary.

methods that accept a function and apply in to all etc. (in parallel).

with array types selectively broken into chunks for better parallel performance.


installation:

	get github.com/splace/tensor3   // 64bit
	//	get github.com/splace/tensor3/32bit
	//	get github.com/splace/tensor3/scaledint
   

Example:  100 x 1 million matrix multiplications, single threaded then parallel.

	package main

	import . "github.com/splace/tensor3"  // 64 bit
	//import . "github.com/splace/tensor3/32bit"
	//import . "github.com/splace/tensor3/scaledint"
	import "time"
	import "fmt"

	func main(){
		ms := make(Matrices, 1000000)
		// fill in array of matrices
		for i := range ms {
			ms[i] = NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
		}
		m := NewMatrix(9, 8, 7, 6, 5, 4, 3, 2, 1)
		start:=time.Now()

		// multiply every matrix by m, do it 100 times (to make benchmark time resolution better.)
		for i := 0; i < 100; i++ {
			ms.Product(m)
		}
		fmt.Println(time.Since(start))
		// same thing in parallel
		Parallel=true
		start=time.Now()
		for i := 0; i < 100; i++ {
			ms.Product(m)
		}
		fmt.Println(time.Since(start))
	}


*/

package tensor3

// Overview/docs: [![GoDoc](https://godoc.org/github.com/splace/tensor3?status.svg)](https://godoc.org/github.com/splace/tensor3)

// Overview/docs:(32bit) [![GoDoc](https://godoc.org/github.com/splace/tensor3/32bit?status.svg)](https://godoc.org/github.com/splace/tensor3/32bit)
