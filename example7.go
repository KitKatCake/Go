package main

import (
	"fmt"
	"time"
	"context"
)


func print2()  {
	fmt.Println("Hello World")
}

func watch(ctx context.Context,name string)  {
	for{
		select {
		case <-ctx.Done():
			fmt.Println(name,"监控退出，停止了...")
			return
		default:
			fmt.Println(name,"goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func main()  {
	ctx,cancel := context.WithCancel(context.Background())

	go watch(ctx,"[监控1]")
	go watch(ctx,"[监控2]")
	go watch(ctx,"[监控3]")

	time.Sleep(time.Second*10)

	fmt.Println("可以了，通知监控停止")
	cancel()
	time.Sleep(5 * time.Second)









	//ch := make(chan int)
	//go func() {
	//	var sum int
	//	for i:=0;i<10;i++{
	//		sum += i
	//	}
	//	ch<-sum
	//}()
	//fmt.Println(<-ch)
	//close(ch)
	//
	//one := make(chan int)
	//two := make(chan int)
	//
	//go func() {
	//	one<-100
	//}()
	//go func() {
	//	v := <-one
	//	two <-v
	//}()
	//
	//fmt.Println(<-two)
	//



	//for{
	//	var i int = 10
	//	if i>0{
	//		go print2()
	//		i--
	//	}
	//}



	for i:=0;i<10;i++{
		//go print2()

	}



	//time.Sleep(time.Second*1)

	//var wg sync.WaitGroup
	//
	//wg.Add(2)
	//
	//go func() {
	//	time.Sleep(2*time.Second)
	//	fmt.Println("1号完成")
	//	wg.Done()
	//
	//}()
	//
	//
	//go func() {
	//	time.Sleep(2*time.Second)
	//	fmt.Println("2号完成")
	//	wg.Done()
	//
	//}()
	//
	//wg.Wait()
	//
	//fmt.Println("好了，大家都干完了，放工")



	//stop := make(chan bool)
	//
	//go func() {
	//	for{
	//		select {
	//		case <-stop:
	//			fmt.Println("监控退出，停止了...")
	//		return
	//		default:
	//			fmt.Println("goroutine监控中...")
	//			time.Sleep(2 * time.Second)
	//		}
	//
	//	}
	//}()
	//
	//time.Sleep(time.Second*10)
	//
	//fmt.Println("可以了，通知监控停止")
	//
	//stop<- true

//	time.Sleep(time.Second*5)




//ctx,cancel := context.WithCancel(context.Background())

//go func(ctx context.Context) {
//	for{
//		select {
//		case <-ctx.Done():
//			fmt.Println("监控退出，停止了...")
//		return
//		default:
//			fmt.Println("goroutine监控中...")
//			time.Sleep(2 * time.Second)
//		}
//
//	}
//}(ctx)
//
//	time.Sleep(10 * time.Second)
//
//	fmt.Println("可以了，通知监控停止")
//	cancel()
//
//	time.Sleep(time.Second * 5)



}


