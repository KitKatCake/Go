package main

import "fmt"

func invoke(a animal2)  {
	a.printInfo2()
}

type animal2 interface {
	printInfo2()
}
type person struct{
	name string
}
type cat2 int

//func (c cat2) printInfo2() {
//	fmt.Println("a cat")
//}
func (c *cat2) printInfo2() {
	fmt.Println("a cat")
}

func(p *person) String()string {
	return "the person name is "+p.name
}

func (p *person) modify(){
	p.name = "李四"
}
func main()  {


	//var c cat2
	//
	//var a animal2
	//
	//a = c
	//
	//a.printInfo2()

	//var c cat

//	invoke(c)

   // var c cat2
    //invoke(&c)

    p := person{name:"张三"}
    //p.modify()
	//(&p).modify()
    //fmt.Println(p.String())
    fmt.Println((&p).String())

}
