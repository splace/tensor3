package main

import . "github.com/splace/tensor3/32bit"
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
	for i := range ms {
		ms[i] = NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
	}
	Parallel=true
	start=time.Now()
	for i := 0; i < 100; i++ {
		ms.Product(m)
	}
	fmt.Println(time.Since(start))	
}

/*  Hal3 Sun 9 Apr 18:53:36 BST 2017 go version go1.6.2 linux/amd64
2.803393529s
2.065988983s
Sun 9 Apr 18:53:42 BST 2017
*/
/*  Hal3 Sun 9 Apr 18:53:49 BST 2017  go version go1.8 linux/amd64

2.833770341s
2.041434037s
Sun 9 Apr 18:53:55 BST 2017
*/
/*  Hal3 Thu 26 Apr 01:49:45 BST 2018  go version go1.10 linux/amd64

2.976279541s
2.705235734s
Thu 26 Apr 01:49:52 BST 2018
*/

