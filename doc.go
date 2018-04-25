/*

Vector, (3-component), and Matrix, (3-Vector).

useful methods on same, operating in-place, not returning reference. (so no single line chained functions.)

arrays of both, called Vectors and Matrices, with their useful methods.

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
			ms[i] = NewMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9)
		}
		m := NewMatrix(9, 8, 7, 6, 5, 4, 3, 2, 1)
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


/*
tensor3.test -test.bench .
As expected: runtime error: index out of range
goos: linux
goarch: arm
BenchmarkMatrixProduct-4         	10000000	       199 ns/op
BenchmarkMatsProduct-4           	  100000	     18915 ns/op
BenchmarkMatsProductParallel-4   	   50000	     33742 ns/op
BenchmarkVecsSum-4               	     100	  12240145 ns/op
BenchmarkVecsSumParallel-4       	     100	  11940562 ns/op
BenchmarkVecsCross-4             	     200	   7850206 ns/op
BenchmarkVecsCrossParallel-4     	     500	   2793741 ns/op
BenchmarkVecsProduct-4           	     100	  11736753 ns/op
BenchmarkVecsProductParallel-4   	     500	   3939572 ns/op
PASS
*/
