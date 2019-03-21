package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}
type ta struct {
	i int
}

//func foo(t ta)  {
//	t.i = 11
//}

func foo(t *ta)  {
	t.i = 11
}
func (c Circle) getArea() float64{
	return math.Pi*c.radius*c.radius
}

func main()  {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))

	var c1 Circle
	c1.radius = 10.00

	fmt.Println("Area of Circle(c1) = ", c1.getArea())

	func(){
		fmt.Println("Hello GO Lanuage!")
	}()

	var a ta
	a.i = 1
	foo(&a)

	fmt.Println(a)




}
