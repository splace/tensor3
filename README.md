/*

**Vector**, (3-**Scalar** x,y,z), and **Matrix**, (3-Vector array) Types.

Useful methods on these, mostly operating in-place (with no returned reference no single line method chaining.)

Arrays of both, (called **Vectors** and **Matrices**), with their useful methods.

Vectorized/SIMD methods that accept a function and apply in to all the array.

selectable single or parallel threaded, transparently, switched with a global var.

**VectorRefs**; array of Vector pointers, with methods as Vectors but also converters to/from Vectors.

the five main types are available (in separate packages) with their base scalar typed as 64bit float, 32bit float, or fixed precision (int scaled for 3dp)



notes:

doesn't use "math" package, left to importers, if necessary.

when parallel, array types are acted on in chunks by different threads.



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

note: there are a lot of methods for common operations on the main types, and documentation comments are often missing on basic methods that appear,to me, unambiguous from their name or i havent got round to yet.

// Overview/docs: [![GoDoc](https://godoc.org/github.com/splace/tensor3?status.svg)](https://godoc.org/github.com/splace/tensor3)

