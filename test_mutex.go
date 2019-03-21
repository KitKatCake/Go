package main

import (
	"fmt"
	"sync/atomic"
)
var value int32
func addValue(delta int32)  {
	for{
		v := atomic.LoadInt32(&value)
		if atomic.CompareAndSwapInt32(&value,v,(v+delta)){
			break
		}
	}
}

func main()  {

	fmt.Println("======old value=======")
	fmt.Println(value)
	fmt.Println("======CAS value=======")
	addValue(3)
	fmt.Println(value)
	fmt.Println("======Store value=======")
	atomic.StoreInt32(&value, 10)
	fmt.Println(value)









	/*mutex := sync.Mutex{}

	cond := sync.NewCond(&mutex)

	for i:=0;i<10;i++{
		go func(num int,cond *sync.Cond) {
			cond.L.Lock()
			if num < 3{
				cond.Wait()
			}
			fmt.Println("num", num, "time", time.Now().Second())
			cond.L.Unlock()
		}(i,cond)
	}

	time.Sleep(time.Second)

	cond.L.Lock()

	cond.Signal()

	cond.L.Unlock()

	time.Sleep(time.Second * 1)

	cond.L.Lock()

	cond.Broadcast()

	cond.L.Unlock()

	time.Sleep(time.Second * 1)

	var i32 int32
	fmt.Println("=====old i32 value=====")
	fmt.Println(i32)

	newI32 := atomic.AddInt32(&i32,3)
	fmt.Println("=====new i32 value=====")
	fmt.Println(i32)
	fmt.Println(newI32)

	var i64 int64
	fmt.Println("=====old i64 value=====")
	fmt.Println(i64)
	newI64 := atomic.AddInt64(&i64,-3)
	fmt.Println("=====new i64 value=====")
	fmt.Println(i64)
	fmt.Println(newI64)

*/

//	fmt.Println(time.Now())
//	fmt.Println(time.Now().Second())
//	fmt.Println(time.Now().Hour())

}
