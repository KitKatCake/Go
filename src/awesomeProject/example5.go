package main

import "fmt"

var c chan int

func getInfo(name string,age int)  {
	fmt.Printf("name: %s,age: %d \n", name, age)
	c <- 0
}
func main()  {

	c = make(chan int)

	go getInfo("li", 12)

	go getInfo("sun", 45)

	fmt.Println("start")

	<-c

	<-c

	close(c)
}