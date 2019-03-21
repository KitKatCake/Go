package main

import (
	"sync"
	"fmt"
	"time"
)
var m *sync.RWMutex
func main() {
	//var wg sync.WaitGroup
	//
	//for i:=0;i<10;i++{
	//	wg.Add(1)
	//	go func(t int) {
	//		defer wg.Done()
	//		fmt.Println(t)
	//	}(i)
	//}
	//wg.Wait()

	wg := sync.WaitGroup{}

	wg.Add(20)

	var rwMutex sync.RWMutex

	Data := 0

	for i := 0; i < 10; i++ {
		go func(t int) {
			rwMutex.RLock()
			defer rwMutex.RUnlock()
			fmt.Printf("Read data: %v\n", Data)
			wg.Done()
			time.Sleep(2 * time.Second)
		}(i)

	go func(t int) {
		rwMutex.Lock()

		defer rwMutex.Unlock()

		Data += t
		fmt.Printf("Write Data: %v %d \n", Data, t)
		wg.Done()

		time.Sleep(2 * time.Second)

	}(i)

}
	time.Sleep(5 * time.Second)
	wg.Wait()
}