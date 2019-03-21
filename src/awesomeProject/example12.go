package main

import (
	"fmt"
	"runtime"
	"sync"
)

func loop(done chan bool)  {
	for i:=0;i<10;i++{
		//runtime.Gosched()
		fmt.Print(i)
	}
	done<-true
}

func main()  {

	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting Go Routines")
	go func() {
		defer wg.Done()

		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()

	go func() {
		defer wg.Done()

		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()



	//runtime.GOMAXPROCS(2)
	//done := make(chan bool)
	//
	//go loop(done)
	//go loop(done)
	//
	//<- done
	//
	//<- done

	//fmt.Println("cpus:", runtime.NumCPU())
    //fmt.Println(runtime.GOMAXPROCS(0))
	//fmt.Println("goroot:", runtime.GOROOT())
	//fmt.Println("os/platform:", runtime.GOOS)
	//
	//name, err := os.Hostname()
	//if err == nil {
	//	fmt.Println(name)
	//}


	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")

}
