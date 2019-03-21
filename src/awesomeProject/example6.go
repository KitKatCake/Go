package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SendInt(ch chan<- int) {
	rand.Seed(time.Now().Unix())
	ch <- rand.Intn(1000)
}

func main()  {
	//var c1, c2, c3 chan int
	//var i1, i2 int

	//i1 = 1
	//i2 = 2

	//ch := make(chan int)
	//
	//go SendInt(ch)
	//  fmt.Print(<-ch)


	intChannels := [3]chan int {
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	rand.Seed(time.Now().Unix())
	index := rand.Intn(3)
	fmt.Printf("The index:%d \n",index)
	intChannels[index] <-index
	select {
	case<-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case<-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case<-intChannels[2]:
		fmt.Println("The third candidate case is selected.")
	default:
		fmt.Println("No candidate case is selected")
	}


	//select {
	//case i1 = <-c1:
	//	fmt.Printf("received ", i1, " from c1\n")
	//case c2 <- i2:
	//	fmt.Printf("sent ", i2, " to c2\n")
	//case i3, ok := (<-c3):  // same as: i3, ok := <-c3
	//	if ok {
	//		fmt.Printf("received ", i3, " from c3\n")
	//	} else {
	//		fmt.Printf("c3 is closed\n")
	//	}
	//default:
	//	fmt.Printf("no communication\n")
	//}
}
