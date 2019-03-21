package main

import "fmt"

func sayHello2(name *string) {
	*name = "Georage"
	fmt.Println("Hello " + *name)

}

func run2(string) string{
     name := "Paul"
     defer sayHello2(&name)

     name = "John"
     return name


}

func main() {
     //name := run2("Tom")
	 //fmt.Println("return: name = " + name)






	var progress = 2
	var target = 8

	title := fmt.Sprintf("已采集%d个药草, 还需要%d个完成任务", progress, target)
	fmt.Println(title)

	pi := 3.14159
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(variant)

	var age = 30
	fmt.Printf("my sge is %d \n",age)

	//b := 10
	var b  = 10
	fmt.Printf("The value of b is %d" ,b)






}
