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
	fmt.Printf("64bit.\t%v\n",time.Since(start)/1000000)	
	Parallel=true
	start=time.Now()
	for i := 0; i < 10; i++ {
		ms.Product(m)
	}
	fmt.Printf("64bit parallel(2 core).\t%v\n",time.Since(start)/1000000)	

	ms32 := make(ttbt.Matrices, 100000)
	for i := range ms {
		ms32[i] = ttbt.NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
	}
	m32 := ttbt.NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
	start=time.Now()
	
	for i := 0; i < 10; i++ {
		ms32.Product(m32)
	}
	fmt.Printf("32bit.\t %v.\n",time.Since(start)/1000000)	
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
	fmt.Printf("Parallel 32bit(2 core)\t %v\n",time.Since(start)/1000000)	
}

/*  Hal3 Sun 9 Apr 18:26:35 BST 2017 go version go1.6.2 linux/amd64
64bit.	41ns
64bit parallel(2 core).	42ns
32bit.	 37ns.
Parallel 32bit(2 core)	 33ns
Sun 9 Apr 18:26:37 BST 2017
*/
/*  Hal3 Sun 9 Apr 18:27:13 BST 2017  go version go1.8 linux/amd64

64bit.	40ns
64bit parallel(2 core).	43ns
32bit.	 28ns.
Parallel 32bit(2 core)	 21ns
Sun 9 Apr 18:27:14 BST 2017
*/

