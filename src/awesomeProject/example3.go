package main

import (
	"fmt"
	"time"
)

func worker(id int,jobs <-chan int,results chan<-int)  {
	for j:=range jobs{
		fmt.Println("worker",id,"processing job",j)
		time.Sleep(time.Second)
		results <- j*2
	}

}

func main()  {
	//jobs := make(chan int,100)
	//results := make(chan int,100)
	//
	//for w:=1;w<=3;w++{
	//	go worker(w,jobs,results)
	//}
	//
	//for j:=1;j<=9;j++{
	//	jobs <-j
	//}
	//close(jobs)
	//
	//for a := 1; a <= 9; a++ {
	//	<-results
	//}

	//var i=3
	//
	//go func(a int) {
	//	fmt.Println(a)
	//	fmt.Println("1")
	//}(i)
	//fmt.Println("2")
	//
	//time.Sleep(time.Second)

	ch :=make(chan int,1)
	//ch <- 1
	go func() {
		ch <-1
        fmt.Print(<-ch)
	}()
	fmt.Println("2")

	time.Sleep(time.Second)


}