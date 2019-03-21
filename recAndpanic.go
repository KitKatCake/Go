package main

import "fmt"

func f1()  {
	fmt.Println("This is f1()")
}

func f2()  {

	defer func() {
		if err := recover();err!=nil{
			fmt.Println("Exception has been caught.")
		}
	}()

	fmt.Println("This is f2()")

	panic(1)
}

func f3()  {
	fmt.Println("This is f3()")
}

func main()  {
	f1()

	f2()

	f3()
}
