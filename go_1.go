package main

import (
	"fmt"
	"time"
)

func gen(nums ...int)<-chan int  {
	out := make(chan int)
	go func() {
		for _,n := range nums{
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n:=range in{
			out <- n * n

		}
		close(out)
	}()
	return out
}

func print3(ch chan int)  {

	ch <- 1

	fmt.Println("Hello World!")

}



func main()  {

	chs := make([]chan int,10)

	for i:=0;i<10;i++{
		chs[i] = make(chan int)

		//defer print3(chs[i])

		go  print3(chs[i])
	}


	for _,ch := range chs{
		<-ch
	}

	time.Sleep(time.Second)



	//fmt.Println("Hello Go!")

	//sigs := make(chan os.Signal,1)
	//done := make(chan bool,1)
	//
	//signal.Notify(sigs,syscall.SIGINT,syscall.SIGTERM)
	//
	//go func() {
	//	sig:= <-sigs
	//	fmt.Println()
	//
	//	fmt.Println(sig)
	//
	//	done <-true
	//}()
	//fmt.Println("awaiting signal")
	//<-done
	//fmt.Println("exiting")


	//c := gen(2,3)
	//
	//out := sq(c)



	//for n:=range sq(sq(gen(2,3))){
	//	fmt.Println(n)
	//}


	//fmt.Println(<-out)

	//fmt.Println(<-out)

}
