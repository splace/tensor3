package main

import . "github.com/splace/tensor3"
import ttbt "github.com/splace/tensor3/32bit"
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

	ttbt.Parallel=false
	ms32 := make(ttbt.Matrices, 1000000)
	for i := range ms {
		ms32[i] = ttbt.NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
	}
	m32 := ttbt.NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
	start=time.Now()
	
	for i := 0; i < 100; i++ {
		ms32.Product(m32)
	}
	fmt.Println(time.Since(start))	
	ttbt.Parallel=true
	start=time.Now()
	for i := 0; i < 100; i++ {
		ms32.Product(m32)
	}
	fmt.Println(time.Since(start))	
}


/*  Hal3 Thu 9 Mar 21:46:29 GMT 2017 go version go1.6.2 linux/amd64
4.071481973s
3.879044162s
2.727359577s
2.081455544s
Thu 9 Mar 21:46:43 GMT 2017
*/

