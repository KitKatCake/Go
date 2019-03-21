package main

import (
	"fmt"
	"github.com/panjf2000/ants"
	"sync"
	"sync/atomic"
	"time"
)


var sum int32
func myFunc(i interface{}) error{
	n := i.(int)
	atomic.AddInt32(&sum,int32(n))
	fmt.Printf("run with %d\n",n)
	return nil
}
func demoFunc() error {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")

	return nil

}

func main()  {

	runTimes := 1000000

	var wg sync.WaitGroup


	for i:=0;i<runTimes;i++{
		wg.Add(1)

		ants.Submit(func() error {
			demoFunc()
			wg.Done()
			return nil
		})
	}
	wg.Wait()

	fmt.Println("running goroutine:%d\n",ants.Running())

	fmt.Println("finish all tasks.\n")

	p,_ := ants.NewPoolWithFunc(10, func(i interface{}) error {
		myFunc(i)
		wg.Done()
		return nil
	})

	for i:=0;i<runTimes;i++{
		wg.Add(1)
		p.Serve(i)
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)

}
