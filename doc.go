/*

Vector, (3-Scalar), and Matrix, (3-Vector) Types.

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
/* benchmark: "Vecs" hal3 Sun 9 Dec 01:22:10 GMT 2018 go version go1.11.2 linux/amd64
goos: linux
goarch: amd64
BenchmarkVecsSum-2               	    2000	    919650 ns/op
BenchmarkVecsSumParallel-2       	    2000	    876863 ns/op
BenchmarkVecsCross-2             	    1000	   1299940 ns/op
BenchmarkVecsCrossParallel-2     	    1000	   1316202 ns/op
BenchmarkVecsProduct-2           	    1000	   1708294 ns/op
BenchmarkVecsProductParallel-2   	    1000	   1750193 ns/op
PASS
ok  	_/run/media/simon/6a5530c2-1442-4e9b-b35f-3db0c9a6984c/home/simon/Dropbox/github/working/tensor3	10.548s
Sun 9 Dec 01:22:22 GMT 2018
*/
/*
goos: windows
goarch: amd64
BenchmarkMatsProduct                	     500	   3248661 ns/op
BenchmarkMatsProduct-2              	     500	   3190975 ns/op
BenchmarkMatsProduct-4              	     500	   3204826 ns/op
BenchmarkMatsProductParallel        	     500	   3224823 ns/op
BenchmarkMatsProductParallel-2      	     500	   2550652 ns/op
BenchmarkMatsProductParallel-4      	    1000	   2320618 ns/op
BenchmarkMatrixProduct              	100000000	        18.2 ns/op
BenchmarkMatrixProduct-2            	100000000	        18.3 ns/op
BenchmarkMatrixProduct-4            	100000000	        18.2 ns/op
BenchmarkVecRefsProduct             	    1000	   1554318 ns/op
BenchmarkVecRefsProduct-2           	    1000	   1496378 ns/op
BenchmarkVecRefsProduct-4           	    1000	   1298330 ns/op
BenchmarkVecRefsProductParallel     	    1000	   1222316 ns/op
BenchmarkVecRefsProductParallel-2   	    2000	    831713 ns/op
BenchmarkVecRefsProductParallel-4   	    2000	    696680 ns/op
BenchmarkVecsSum                    	    2000	    887726 ns/op
BenchmarkVecsSum-2                  	    2000	    914732 ns/op
BenchmarkVecsSum-4                  	    2000	    908229 ns/op
BenchmarkVecsSumParallel            	    2000	    905729 ns/op
BenchmarkVecsSumParallel-2          	    3000	    575506 ns/op
BenchmarkVecsSumParallel-4          	    5000	    399101 ns/op
BenchmarkVecsCross                  	    2000	    924775 ns/op
BenchmarkVecsCross-2                	    2000	    904292 ns/op
BenchmarkVecsCross-4                	    2000	    896726 ns/op
BenchmarkVecsCrossParallel          	    2000	    949780 ns/op
BenchmarkVecsCrossParallel-2        	    2000	    690675 ns/op
BenchmarkVecsCrossParallel-4        	    3000	    530802 ns/op
BenchmarkVecsProduct                	    1000	   1229315 ns/op
BenchmarkVecsProduct-2              	    2000	   1155794 ns/op
BenchmarkVecsProduct-4              	    2000	   1172296 ns/op
BenchmarkVecsProductParallel        	    1000	   1238307 ns/op
BenchmarkVecsProductParallel-2      	    2000	    888688 ns/op
BenchmarkVecsProductParallel-4      	    2000	    687639 ns/op
PASS
*/

/* benchmark: "." hal3 Mon 31 Dec 22:31:36 GMT 2018 go version go1.11.4 linux/amd64
goos: linux
goarch: amd64
BenchmarkMatrixProduct              	50000000	        24.1 ns/op
BenchmarkMatrixProduct-2            	50000000	        24.3 ns/op
BenchmarkMatsProduct                	     300	   4513350 ns/op
BenchmarkMatsProduct-2              	     300	   4520161 ns/op
BenchmarkMatsProductParallel        	     300	   4528445 ns/op
BenchmarkMatsProductParallel-2      	     300	   4678347 ns/op
BenchmarkVecRefsProduct             	    1000	   2161834 ns/op
BenchmarkVecRefsProduct-2           	    1000	   2137820 ns/op
BenchmarkVecRefsProductParallel     	    1000	   2228517 ns/op
BenchmarkVecRefsProductParallel-2   	    1000	   2189679 ns/op
BenchmarkVecsSum                    	    2000	    955081 ns/op
BenchmarkVecsSum-2                  	    2000	    951417 ns/op
BenchmarkVecsSumParallel            	    2000	    975744 ns/op
BenchmarkVecsSumParallel-2          	    2000	    885654 ns/op
BenchmarkVecsCross                  	    1000	   1333085 ns/op
BenchmarkVecsCross-2                	    1000	   1343820 ns/op
BenchmarkVecsCrossParallel          	   20000	     93123 ns/op
BenchmarkVecsCrossParallel-2        	   10000	    106036 ns/op
BenchmarkVecsProduct/100            	 1000000	      1617 ns/op
BenchmarkVecsProduct/100-2          	 1000000	      1620 ns/op
BenchmarkVecsProduct/10000          	   10000	    162212 ns/op
BenchmarkVecsProduct/10000-2        	   10000	    164243 ns/op
BenchmarkVecsProduct/40000000       	       2	 710262528 ns/op
BenchmarkVecsProduct/40000000-2     	       2	 706344689 ns/op
BenchmarkVecsProduct/100_Parallel   	  500000	      4071 ns/op
BenchmarkVecsProduct/100_Parallel-2 	  500000	      4059 ns/op
BenchmarkVecsProduct/10000_Parallel           	   10000	    162319 ns/op
BenchmarkVecsProduct/10000_Parallel-2         	   10000	    176854 ns/op
BenchmarkVecsProduct/40000000_Parallel        	       2	 756703664 ns/op
BenchmarkVecsProduct/40000000_Parallel-2      	       2	 656509871 ns/op
BenchmarkVecsFindNearest/100                  	 2000000	       958 ns/op
BenchmarkVecsFindNearest/100-2                	 2000000	       924 ns/op
BenchmarkVecsFindNearest/10000                	   20000	     90606 ns/op
BenchmarkVecsFindNearest/10000-2              	   20000	     90877 ns/op
BenchmarkVecsFindNearest/40000000             	       3	 403270754 ns/op
BenchmarkVecsFindNearest/40000000-2           	       3	 404207765 ns/op
BenchmarkVecsFindNearest/100_Parallel         	 2000000	       925 ns/op
BenchmarkVecsFindNearest/100_Parallel-2       	 2000000	       922 ns/op
BenchmarkVecsFindNearest/10000_Parallel       	   20000	     92883 ns/op
BenchmarkVecsFindNearest/10000_Parallel-2     	   10000	    107317 ns/op
BenchmarkVecsFindNearest/40000000_Parallel    	       3	 436411858 ns/op
BenchmarkVecsFindNearest/40000000_Parallel-2  	       5	 327502446 ns/op
PASS
ok  	_/run/media/simon/6a5530c2-1442-4e9b-b35f-3db0c9a6984c/home/simon/Dropbox/github/working/tensor3	93.765s
Mon 31 Dec 22:33:11 GMT 2018
*/

