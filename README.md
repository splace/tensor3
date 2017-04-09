/*

Vector, (3-component), and Matrix, (3-Vector) Types.

useful methods on same, operating in-place, not returning reference. (so no single line chained functions.)

arrays of both, called Vectors and Matrices, with their useful methods.

VectorRefs, array of pointers to Vector, with methods to convert to/from Vectors.

64 or 32bit

operations can be switched, with a global var, between single and parallel thread solutions.

(import package multiple times, with different aliases, to enable simultaneous use of both modes.)

arrays selectively broken into chunks for optimised parallelization.


installation:

	get github.com/splace/tensor3
	get github.com/splace/tensor3/32bit



Example:  100 x 1 million matrix multiplications, single threaded then parallel.

	package main

	import . "github.com/splace/tensor3"
	import "time"
	import "fmt"

	func main(){
		ms := make(Matrices, 1000000)
		for i := range ms {
			ms[i] = NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
		}
		m := NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
		start:=time.Now()

		for i := 0; i < 100; i++ {
			ms.Product(m)
		}
		fmt.Println(time.Since(start))
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
