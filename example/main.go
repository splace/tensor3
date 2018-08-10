package main

import . "github.com/splace/tensor3"
import ttbt "github.com/splace/tensor3/32bit"
import sit "github.com/splace/tensor3/scaledint"
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
	for i := range ms32 {
		ms32[i] = ttbt.NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
	}
	m32 := ttbt.NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
	start=time.Now()
	
	for i := 0; i < repts; i++ {
		ms32.Product(m32)
	}
	t=time.Since(start)
	fmt.Printf("32bit.\t %v.\n",t/size/repts)	
	ttbt.Parallel=true
	start=time.Now()
	for i := 0; i < repts; i++ {
		ms32.Product(m32)
	}
	t=time.Since(start)
	fmt.Printf("Parallel 32bit(%d core)\t %v\n",runtime.NumCPU(),t/size/repts)	

	mssi := make(sit.Matrices, size)
	for i := range mssi {
		mssi[i] = sit.NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
	}
	msi := sit.NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
	start=time.Now()
	for i := 0; i < repts; i++ {
		mssi.Product(msi)
	}
	t=time.Since(start)
	fmt.Printf("int\t %v\n",t/size/repts)	

	sit.Parallel=true
	start=time.Now()
	for i := 0; i < repts; i++ {
		mssi.Product(msi)
	}
	t=time.Since(start)
	fmt.Printf("Parallel int(%d core)\t %v\n",runtime.NumCPU(),t/size/repts)	

}
/*  Hal3 Thu 26 Apr 01:44:18 BST 2018 go version go1.6.2 linux/amd64
64bit.	41ns
64bit parallel(2 core).	40ns
32bit.	 27ns.
Parallel 32bit(2 core)	 27ns
int(2 core)	 52ns
Parallel int(2 core)	 52ns
Thu 26 Apr 01:44:44 BST 2018
*/
/*  Hal3 Thu 26 Apr 01:44:55 BST 2018  go version go1.10 linux/amd64

64bit.	40ns
64bit parallel(2 core).	40ns
32bit.	 28ns.
Parallel 32bit(2 core)	 27ns
int(2 core)	 43ns
Parallel int(2 core)	 43ns
Thu 26 Apr 01:45:26 BST 2018
*/
/* run: tags="" hal3 Thu 9 Aug 18:19:06 BST 2018 go version go1.10.3 linux/amd64
64bit.	43ns
64bit parallel(2 core).	43ns
32bit.	 29ns.
Parallel 32bit(2 core)	 27ns
int(2 core)	 45ns
Parallel int(2 core)	 45ns
Thu 9 Aug 18:19:33 BST 2018
*/

