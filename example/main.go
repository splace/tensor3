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


/*  Hal3 Thu 9 Mar 21:46:29 GMT 2017 go version go1.6.2 linux/amd64
4.071481973s
3.879044162s
2.727359577s
2.081455544s
Thu 9 Mar 21:46:43 GMT 2017
*/
/*  Hal3 Thu 9 Mar 22:13:54 GMT 2017  go version go1.8 linux/amd64

4.025525437s
3.875234189s
2.889553996s
2.010033329s
Thu 9 Mar 22:14:11 GMT 2017
*/
/*  Hal3 Thu 9 Mar 23:25:55 GMT 2017 go version go1.6.2 linux/amd64
{0 0 0}
4.070981192s
3.868496532s
2.742058063s
2.075840841s
Thu 9 Mar 23:26:09 GMT 2017
*/
/*  Hal3 Thu 9 Mar 23:26:30 GMT 2017 go version go1.6.2 linux/amd64
0 0 0
4.281453565s
4.106664483s
2.737290382s
2.129397773s
Thu 9 Mar 23:26:44 GMT 2017
*/

