package main

func generate(ch chan<- int)  {

	for i:=2;;i++{
		ch <- i
	}
}

func filter(in <-chan int, out chan <- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main()  {

	ch :=make(chan int)
	go generate(ch)

	for i:=0;i<10;i++{
		prime := <-ch
		print(prime,"n")

		ch1 := make(chan int)
		go filter(ch,ch1,prime)

		ch = ch1

	}


	//requests := make(chan int,5)
	//
	//for i:=1;i<=5;i++{
	//	requests <- i
	//}
	//close(requests)
	//
	//limiter := time.Tick(200*time.Millisecond)
	//
	//for req := range requests{
	//	<-limiter
	//	fmt.Println("request",req,time.Now())
	//}
	//
	//burstyLimiter := make(chan time.Time,3)
	//
	//for i:=0;i<3;i++{
	//	burstyLimiter <- time.Now()
	//}
	//
	//go func() {
	//	for t:= range time.Tick(200*time.Millisecond){
	//		burstyLimiter <- t
	//	}
	//}()
	//
	//burstyRequests := make(chan int, 5)
	//for i := 1; i <= 5; i++ {
	//	burstyRequests <- i
	//}
	//close(burstyRequests)
	//for req := range burstyRequests {
	//	<-burstyLimiter
	//	fmt.Println("request", req, time.Now())
	//}


	//runtime.GOMAXPROCS(1)
	//ch := make(chan int)
	//go func(){
	//	fmt.Printf("开始阻塞1。。。")
	//	time.Sleep(time.Second)
	//	fmt.Printf("结束阻塞1。。。")
	//	ch <- 4
	//}()
	//
	//fmt.Printf("ch ==>%d", <- ch)

}
