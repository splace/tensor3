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

*/

package tensor3


// some benchmarking

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

/* benchmark: "" hal3 Sat 28 Apr 20:56:16 BST 2018 go version go1.10 linux/amd64
goos: linux
goarch: amd64
BenchmarkMatrixProduct-2            	50000000	        26.3 ns/op
BenchmarkMatsProduct-2              	     300	   4716321 ns/op
BenchmarkMatsProductParallel-2      	     300	   4758117 ns/op
BenchmarkVecRefsProduct-2           	     500	   2324318 ns/op
BenchmarkVecRefsProductParallel-2   	    1000	   2418231 ns/op
BenchmarkVecsSum-2                  	    2000	    890864 ns/op
BenchmarkVecsSumParallel-2          	    2000	    894359 ns/op
BenchmarkVecsCross-2                	    1000	   1425029 ns/op
BenchmarkVecsCrossParallel-2        	    1000	   1456007 ns/op
BenchmarkVecsProduct-2              	    1000	   1989979 ns/op
BenchmarkVecsProductParallel-2      	    1000	   1980549 ns/op
PASS
ok  	
Sat 28 Apr 20:56:37 BST 2018
*/

/*
tensor3.test -test.bench Vecs
As expected: runtime error: index out of range
goos: linux
goarch: arm
BenchmarkVecsSum-4               	     100	  12240145 ns/op
BenchmarkVecsSumParallel-4       	     100	  11940562 ns/op
BenchmarkVecsCross-4             	     200	   7850206 ns/op
BenchmarkVecsCrossParallel-4     	     500	   2793741 ns/op
BenchmarkVecsProduct-4           	     100	  11736753 ns/op
BenchmarkVecsProductParallel-4   	     500	   3939572 ns/op
PASS
*/
/*  Hal3 Wed 25 Apr 21:44:55 BST 2018  go version go1.10 linux/amd64

goos: linux
goarch: amd64
BenchmarkVecsSum-2               	    2000	    838379 ns/op
BenchmarkVecsSumParallel-2       	    2000	    843700 ns/op
BenchmarkVecsCross-2             	    1000	   1224486 ns/op
BenchmarkVecsCrossParallel-2     	    1000	   1160309 ns/op
BenchmarkVecsProduct-2           	    1000	   1882070 ns/op
BenchmarkVecsProductParallel-2   	    1000	   1418380 ns/op
PASS
ok  	
Wed 25 Apr 21:45:05 BST 2018
*/

/* benchmark: "" hal3 Sat 28 Apr 21:10:56 BST 2018 go version go1.10 linux/amd64
goos: linux
goarch: amd64
BenchmarkMatrixProduct-2    		   	50000000	        25.4 ns/op
BenchmarkMatsProduct-2          	 	     300	   4034869 ns/op
BenchmarkMatsProductParallel-2   		     500	   4047186 ns/op
BenchmarkVecRefsProduct-2           	    1000	   2109461 ns/op
BenchmarkVecRefsProductParallel-2   	    1000	   1858104 ns/op
BenchmarkVecsSum-2                  	    2000	    858996 ns/op
BenchmarkVecsSumParallel-2          	    3000	    548285 ns/op
BenchmarkVecsCross-2                	    2000	   1189825 ns/op
BenchmarkVecsCrossParallel-2        	    2000	   1041651 ns/op
BenchmarkVecsProduct-2              	    1000	   1850390 ns/op
BenchmarkVecsProductParallel-2      	    1000	   1193208 ns/op
PASS
ok  	
Sat 28 Apr 21:11:19 BST 2018
*/
/* benchmark: "Vecs" hal3 Sat 19 May 20:51:24 BST 2018 go version go1.10.2 linux/amd64
goos: linux
goarch: amd64
BenchmarkVecsSum-2               	    2000	    898831 ns/op
BenchmarkVecsSumParallel-2       	    2000	    943361 ns/op
BenchmarkVecsCross-2             	    1000	   1281549 ns/op
BenchmarkVecsCrossParallel-2     	    1000	   1327454 ns/op
BenchmarkVecsProduct-2           	    1000	   1888471 ns/op
BenchmarkVecsProductParallel-2   	    1000	   1944907 ns/op
PASS
ok  	_/run/media/simon/6a5530c2-1442-4e9b-b35f-3db0c9a6984c/home/simon/Dropbox/github/working/tensor3	11.104s
Sat 19 May 20:51:36 BST 2018
*/


/* benchmark: "Vecs" hal3 Wed 16 May 23:29:56 BST 2018 go version go1.10 linux/amd64
goos: linux
goarch: amd64
BenchmarkVecsSum-2               	    1000	   1511522 ns/op
BenchmarkVecsSumParallel-2       	    1000	   1094181 ns/op
BenchmarkVecsCross-2             	     500	   2372967 ns/op
BenchmarkVecsCrossParallel-2     	    1000	   1639160 ns/op
BenchmarkVecsProduct-2           	     500	   2421282 ns/op
BenchmarkVecsProductParallel-2   	     500	   2247074 ns/op
PASS
ok  	_/run/media/simon/6a5530c2-1442-4e9b-b35f-3db0c9a6984c/home/simon/Dropbox/github/working/tensor3	9.206s
Wed 16 May 23:30:06 BST 2018
*/
/*  Hal3 Sat 19 May 21:45:09 BST 2018  go version go1.10 linux/amd64

goos: linux
goarch: amd64
BenchmarkVecsSum-2               	    2000	    875164 ns/op
BenchmarkVecsSumParallel-2       	    3000	    558054 ns/op
BenchmarkVecsCross-2             	    1000	   1269314 ns/op
BenchmarkVecsCrossParallel-2     	    2000	   1126794 ns/op
BenchmarkVecsProduct-2           	    1000	   1862880 ns/op
BenchmarkVecsProductParallel-2   	    1000	   1222321 ns/op
PASS
ok  	_/home/simon/Dropbox/github/working/tensor3	10.834s
Sat 19 May 21:45:20 BST 2018
*/

