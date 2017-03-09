package main

import . "../../tensor3"
import "time"
import "fmt"

func main(){
	ms := make(Matrixes, 1000000)
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
}
/*  Hal3 Tue 7 Mar 00:58:11 GMT 2017 go version go1.6.2 linux/amd64
Tue 7 Mar 00:58:11 GMT 2017
*/
/*  Hal3 Tue 7 Mar 00:58:18 GMT 2017 go version go1.6.2 linux/amd64
Tue 7 Mar 00:58:18 GMT 2017
*/
/*  Hal3 Tue 7 Mar 00:58:26 GMT 2017 go version go1.6.2 linux/amd64
Tue 7 Mar 00:58:26 GMT 2017
*/
/*  Hal3 Tue 7 Mar 00:59:58 GMT 2017 go version go1.6.2 linux/amd64
4.131088887s
Tue 7 Mar 01:00:03 GMT 2017
*/
/*  Hal3 Tue 7 Mar 01:00:29 GMT 2017 go version go1.6.2 linux/amd64
Tue 7 Mar 01:00:29 GMT 2017
*/
/*  Hal3 Tue 7 Mar 01:01:09 GMT 2017 go version go1.6.2 linux/amd64
5.164817621s
5.003088886s
Tue 7 Mar 01:01:20 GMT 2017
*/

