package main

import (
	"fmt"
)

func Greeting(prefix string,who ... string)  {
	fmt.Println(prefix)

	for _,name := range who{
		fmt.Println(name)
	}

}

func run()(string)  {
	name := "Paul"
	aFun := func() {
		fmt.Println("Hello "+name)
		name = "Georage"
		fmt.Println("Hello "+name)
	}
	name = "John"
	aFun()
	return name
}

func main()  {

	//name := run()

	//fmt.Println("return name= "+name)

	Greeting("Hello:", "tom", "mike", "jesse")

	var i *int
	i=new(int)
	*i=10
	fmt.Println(*i)







}
