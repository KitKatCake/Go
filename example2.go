package main

import (
	"fmt"
)

type V struct {
	i int32
	j int64
}

func (this V)PutI() {
	fmt.Printf("i=%d\n", this.i)
}

func (this V) PutJ() {
	fmt.Printf("j=%d\n", this.j)
}

func foo(fm map[string]int)  {
	fm["a"] = 11
}

//var gm map[int]int

var gm = make(map[int]int)


func foo2(i int)  {
	for j := 0; j < 10; j++ {
		gm[i] = j
	}
	for i,v := range gm{
		fmt.Println(i,v)
	}

}

func main()  {
	//var v *V = new(V)
	//var i *int32 = (*int32)(unsafe.Pointer(v))

	//m :=make(map[string]int)
	//m["a"] = 1
	//fmt.Println(m)
	//
	//foo(m)
	//fmt.Println(m)


	//var mutex sync.Mutex

	//gm = make(map[int]int)

	//mutex.Lock()

	for i:=0;i<10;i++{
	//	go foo2(i)
		foo2(i)
	}

	//mutex.Unlock()
	//time.Sleep(2*time.Second)







}
