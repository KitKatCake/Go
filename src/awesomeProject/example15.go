package main

import (
	"fmt"
	//"time"
)

//func worker(id int,jobs <-chan int,results chan <-int)  {
//	for j:=range jobs{
//
//	}
//}

func intSeq() func() int{
	i :=0
	return func() int {
		i++
		return i
	}


}

func main()  {


	//c1 := make(chan string)
	//c2 := make(chan string)
	//
	//go func() {
	//	time.Sleep(1*time.Second)
	//	c1 <- "one"
	//}()
	//
	//go func() {
	//	time.Sleep(2*time.Second)
	//	c2 <- "two"
	//}()
	//
	//
	//for i:=0;i<2;i++{
	//	select {
	//	case msg1:= <-c1:
	//		fmt.Println("received",msg1)
	//	case msg2:= <- c2:
	//		fmt.Println("received",msg2)
	//	}
	//
	//
	//}



	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())


}




