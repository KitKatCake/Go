package main

import "fmt"

func main()  {
	//a := func() int{
	//	var i int = 5
	//	fmt.Println("func 1")
	//	return i
	//}()

	//fmt.Println(a)

	//func(arg int){
	//	fmt.Printf("func %d\n",arg)
	//}(2)



	//b := func() int{
	//	fmt.Printf("func3\n")
	//	return 3
	//}
	//
	//fmt.Println(b())
	j := 5
	c := func()(func ()) {
		i:=10
		return func() {
			fmt.Printf("i,j:%d,%d\n",i,j)
		}
	}()

	c()

	j *= 2

	c()




}
