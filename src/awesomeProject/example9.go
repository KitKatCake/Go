package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:",s)
	return s
}

func un(s string)  {
	fmt.Println("leaving:",s)
}
func a()  {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b()  {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func run()  {
	name := "Paul"
	defer sayHello(name)
	name = "John"
}

func sayHello(name string)  {
	fmt.Println("Hello "+name)
}

func main()  {
	//for i:=0;i<5;i++{
	//	defer fmt.Println("%d",i)
	//	fmt.Println("bbbbb")
	//}
	//fmt.Println("aaaaa")


	//b()

	run()

	fmt.Println("Ha Ha Ha Ha!!!")
}
