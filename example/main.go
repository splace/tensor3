package main

import . "../../tensor3"
import ttbt "../../tensor3/32bit"
import ttbtp "../../tensor3/32bit"
import "time"
import "fmt"

func main(){
	ms := make(Matrices, 100000)
	for i := range ms {
		ms[i] = NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
	}
	m := NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
	start:=time.Now()
	
	for i := 0; i < 10; i++ {
		ms.Product(m)
	}
	fmt.Println(time.Since(start))	
	Parallel=true
	start=time.Now()
	for i := 0; i < 10; i++ {
		ms.Product(m)
	}
	fmt.Println(time.Since(start))	

	ms32 := make(ttbt.Matrices, 100000)
	for i := range ms {
		ms32[i] = ttbt.NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
	}
	m32 := ttbt.NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
	start=time.Now()
	
	for i := 0; i < 10; i++ {
		ms32.Product(m32)
	}
	fmt.Println(time.Since(start))	
	ms32p := make(ttbtp.Matrices, 100000)
	for i := range ms {
		ms32[i] = ttbtp.NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
	}
	m32p := ttbtp.NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
	ttbtp.Parallel=true
	start=time.Now()
	for i := 0; i < 10; i++ {
		ms32p.Product(m32p)
	}
	fmt.Println(time.Since(start))	
}

/*  Hal3 Fri 10 Mar 16:27:07 GMT 2017 go version go1.6.2 linux/amd64
Fri 10 Mar 16:27:07 GMT 2017
*/
/*  Hal3 Fri 10 Mar 16:27:30 GMT 2017 go version go1.6.2 linux/amd64
41.722678ms
43.37056ms
27.509332ms
28.380486ms
Fri 10 Mar 16:27:31 GMT 2017
*/

