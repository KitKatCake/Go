package main

import (
	"fmt"
	"sync"
)

func say(s string)  {
	for i:=0;i<2;i++{
		//runtime.Gosched()

		fmt.Println(s)
	}

}


func main()  {
    //runtime.GOMAXPROCS(2)
	//go say("world")
	//say("hello")
	//
	//time.Sleep(time.Second)

	var l *sync.Mutex
	l = new(sync.Mutex)
	l.Lock()
	defer l.Unlock()
	fmt.Println("1")

}
