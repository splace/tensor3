package main

import . "../../tensor3"
import ttbt "../../tensor3/32bit"
import ttbtp "../../tensor3/32bit"
import sit "../../tensor3/scaledint"
import sitp "../../tensor3/scaledint"
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
	
	ms32p := make(ttbtp.Matrices, size)
	for i := range ms32p {
		ms32p[i] = ttbtp.NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
	}
	m32p := ttbtp.NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
	ttbtp.Parallel=true
	start=time.Now()
	for i := 0; i < repts; i++ {
		ms32p.Product(m32p)
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
	fmt.Printf("int(%d core)\t %v\n",runtime.NumCPU(),t/size/repts)	

	mssip := make(sitp.Matrices, size)
	for i := range mssip {
		mssip[i] = sitp.NewMatrix(1, 2, 3,4, 5, 6,7, 8, 9)
	}
	msip := sitp.NewMatrix(9, 8, 7,6, 5, 4,3, 2, 1)
	sitp.Parallel=true
	start=time.Now()
	for i := 0; i < repts; i++ {
		mssip.Product(msip)
	}
	t=time.Since(start)
	fmt.Printf("Parallel int(%d core)\t %v\n",runtime.NumCPU(),t/size/repts)	

}

/*  Hal3 Wed 10 May 21:28:51 BST 2017  go version go1.8 linux/amd64

64bit.	42ns
64bit parallel(2 core).	42ns
32bit.	 25ns.
Parallel 32bit(2 core)	 19ns
Wed 10 May 21:29:06 BST 2017
*/
/*  Hal3 Sun 11 Jun 19:18:05 BST 2017 go version go1.6.2 linux/amd64
64bit.	44ns
64bit parallel(2 core).	39ns
32bit.	 28ns.
Parallel 32bit(2 core)	 20ns
int(2 core)	 57ns
Parallel int(2 core)	 40ns
Sun 11 Jun 19:18:36 BST 2017
*/
/*  Hal3 Sun 11 Jun 19:20:08 BST 2017 go version go1.6.2 linux/amd64
64bit.	45ns
64bit parallel(2 core).	42ns
32bit.	 28ns.
Parallel 32bit(2 core)	 23ns
int(2 core)	 59ns
Parallel int(2 core)	 44ns
Sun 11 Jun 19:20:35 BST 2017
*/
/*  Hal3 Sun 11 Jun 19:21:34 BST 2017  go version go1.8.3 linux/amd64

64bit.	45ns
64bit parallel(2 core).	42ns
32bit.	 32ns.
Parallel 32bit(2 core)	 23ns
int(2 core)	 55ns
Parallel int(2 core)	 46ns
Sun 11 Jun 19:22:05 BST 2017
*/
/*  Hal3 Fri 29 Sep 23:36:11 BST 2017 go version go1.6.2 linux/amd64
64bit.	41ns
64bit parallel(2 core).	38ns
32bit.	 27ns.
Parallel 32bit(2 core)	 19ns
int(2 core)	 57ns
Parallel int(2 core)	 39ns
Fri 29 Sep 23:36:44 BST 2017
*/
/*  Hal3 Fri 29 Sep 23:37:12 BST 2017  go version go1.9 linux/amd64

64bit.	41ns
64bit parallel(2 core).	39ns
32bit.	 29ns.
Parallel 32bit(2 core)	 20ns
int(2 core)	 51ns
Parallel int(2 core)	 40ns
Fri 29 Sep 23:37:43 BST 2017
*/

