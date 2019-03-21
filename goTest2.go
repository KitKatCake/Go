package main

import (
	"bytes"
	"fmt"
	"io"
)
type animal interface {
	printInfo()
}
type cat int
type dog int

func (c cat) printInfo() {
fmt.Println("a cat")
}

func (d dog) printInfo()  {
	fmt.Println("a dog")
}
func main()  {
	var b bytes.Buffer
	fmt.Fprint(&b,"Hello World")
	//fmt.Println(b.String())
	var w io.Writer

	w = &b

	fmt.Println(w)


	var a animal

	var c cat

	a = c

	a.printInfo()

	var d dog

	a = d

	a.printInfo()
}
