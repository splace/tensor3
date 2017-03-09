package main

import . "../../tensor3"
import ttbt "../../tensor3/32bit"
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


