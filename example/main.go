package main

import . "../../tensor3"
import ttbt "../../tensor3/32bit"
import ttbtp "../../tensor3/32bit"
import "time"
import "fmt"
import "runtime"


const size = 1000000
const repts=100

func main(){
	ms := make(Matrices, size)
	for i := range ms {
		ms[i] = NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
	}
	m := NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
	start:=time.Now()
	
	for i := 0; i < repts; i++ {
		ms.Product(m)
	}
	t:=time.Since(start)
	fmt.Printf("64bit.\t%v\n",t/size/repts)	
	Parallel=true
	start=time.Now()
	for i := 0; i < repts; i++ {
		ms.Product(m)
	}
	t=time.Since(start)
	fmt.Printf("64bit parallel(%d core).\t%v\n",runtime.NumCPU(),t/size/repts)	

	ms32 := make(ttbt.Matrices, size)
	for i := range ms {
		ms32[i] = ttbt.NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
	}
	m32 := ttbt.NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
	start=time.Now()
	
	for i := 0; i < repts; i++ {
		ms32.Product(m32)
	}
	t=time.Since(start)
	fmt.Printf("32bit.\t %v.\n",t/size/repts)	
	ms32p := make(ttbtp.Matrices, size)
	for i := range ms {
		ms32[i] = ttbtp.NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
	}
	m32p := ttbtp.NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
	ttbtp.Parallel=true
	start=time.Now()
	for i := 0; i < repts; i++ {
		ms32p.Product(m32p)
	}
	t=time.Since(start)
	fmt.Printf("Parallel 32bit(%d core)\t %v\n",runtime.NumCPU(),t/size/repts)	
}

/*  Hal3 Wed 10 May 21:28:51 BST 2017  go version go1.8 linux/amd64

64bit.	42ns
64bit parallel(2 core).	42ns
32bit.	 25ns.
Parallel 32bit(2 core)	 19ns
Wed 10 May 21:29:06 BST 2017
*/

