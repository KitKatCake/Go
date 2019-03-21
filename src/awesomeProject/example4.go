package main

import (
	"fmt"
	"time"
)

func produce(p chan <-int)  {
	for i:=0;i<10;i++{
		p<-i
		fmt.Println("send: ",i)
	}
}

func consumer(c <-chan int)  {
	for i:=0;i<10;i++{
		fmt.Println("receive:", <-c)
	}
}


func main()  {
	//ch := make(chan int)
	ch := make(chan int,10)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}
