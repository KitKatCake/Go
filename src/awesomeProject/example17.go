package main

import "fmt"
import "github.com/panjf2000/ants"

func main()  {

	resule := getNum()

	i := 10

	fmt.Println(resule(i))

	//func(){
	//fmt.Println("Hello Go Lanuage")
	//}()


	//myfunc := exfunc()
	//
	//for i:=0;i<10;i++{
	//	fmt.Printf("%10d\n",myfunc(i))
	//}
}

func exfunc()  func(int) int{
	sum :=0
	return func(x int) int {
		sum += x
		return sum
	}
}

func getNum() func(int) int {
	//i :=0
	return func(i int) int {
		i += 1
		return i
	}
}


