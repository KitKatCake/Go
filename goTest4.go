package main

import (
	"fmt"
	"sync"
	"runtime"
)

func print(a ...interface{})  {
	for _,v :=range a{
		fmt.Println(v)
	}
}
func main()  {
	//runtime.GOMAXPROCS(4)
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i:=1;i<100;i++{
			fmt.Println("A:",i)
		}
	}()

	go func() {
		defer wg.Done()
		for i:=1;i<100;i++{
			fmt.Println("B:",i)
		}
	}()

	wg.Wait()


	//print("1","2","3")
	//log.SetFlags(log.Ldate|log.Lshortfile)
	//log.Println("飞雪无情的博客:","http://www.flysnow.org")
	//log.Printf("飞雪无情的微信公众号：%s\n","flysnow_org")

}
