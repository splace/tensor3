/*

Vector, (3-component), and Matrix, (3-Vector) Types.

Useful methods on same, operating in-place, not returning reference. (so no single line chained functions.)

Arrays, called Vectors and Matrices, with their useful methods (parallel).

VectorRefs, array of pointers to Vector, with methods to convert to/from Vectors.

64bit float, 32bit float, or fixed scale int component value types (separate packages).

(import package multiple times, with different aliases, to enable simultaneous use.)

doesn't use "math" package, left to importers, if necessary. (callbacks to use parallel functionality).

single and parallel thread solutions, can be switched, with a global var.

with array types selectively broken into chunks for better parallel performance.


installation:

	get github.com/splace/tensor3    

Example:  100 x 1 million matrix multiplications, single threaded then parallel.

	package main

	import . "github.com/splace/tensor3"
	import "time"
	import "fmt"

	func main(){
		ms := make(Matrices, 1000000)
		// fill in array of matrices
		for i := range ms {
			ms[i] = NewMatrix(1, 2, 3,4, 5, 6, 7, 8, 9)
		}
		m := NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
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
