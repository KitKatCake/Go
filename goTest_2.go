package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())

	testCond()

}

func testCond()  {
	c := sync.NewCond(&sync.Mutex{})

	condition := false

	go func() {

		time.Sleep(time.Second)

		c.L.Lock()

		fmt.Println("[1] 变更condition状态,并发出变更通知.")

		condition = true

		c.Signal()

		fmt.Println("[1] 继续后续处理.")

		c.L.Unlock()
	}()

	c.L.Lock()

	fmt.Println("[2] condition..........1")

	for !condition{
		fmt.Println("[2] condition..........2")
		c.Wait()
		fmt.Println("[2] condition..........3")
	}

	fmt.Println("[2] condition..........4")
	c.L.Unlock()

	fmt.Println("main end...")
































}
