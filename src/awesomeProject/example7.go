package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Generate_Randnum() int{
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(100)

	fmt.Println("rand is %v",rnd)

	return rnd

}

type Element interface {

}

func getSequence() func() int {
	i :=0
	return func() int {
		i +=1
		return i
	}
}
func add(x1,x2 int)func()(int,int)  {
	i :=0
	return func() (int, int) {
		i++
		return i,x1+x2
	}
}

func add2(x1,x2 int) func(x3,x4 int)(int,int,int) {
	i := 0
	return func(x3 int,x4 int) (int,int,int){
		i++
		return i,x1+x2,x3+x4
	}
}

func main()  {


	//Generate_Randnum()

	s := "BrainWu"
	if v, ok := interface{}(s).(string); ok {
		fmt.Println(v)
	}

	//var a int = 1
	//
	//b := int64(a)
	//
	//fmt.Println(b)
	//
	//var e Element = 100
	//var e Element = "string"
	var e Element = 200.369
	switch value:=e.(type) {
	case int:
		fmt.Println("int",value)
	case string:
		fmt.Println("string",value)
	default:
		fmt.Println("unknown",value)

	}

	//fmt.Println(rand.Int())
	//fmt.Println(rand.Int())


	add_func := add(1,2)
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())

    fmt.Println("-------------------------------------------")

	add_func2 := add2(1,2)
	fmt.Println(add_func2(1,1))
	fmt.Println(add_func2(0,0))
	fmt.Println(add_func2(2,2))

	fmt.Println("-------------------------------------------")


	nextNumber :=getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())


	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())


	//var sum int = 17
	//var count int = 5
	//var mean float32
	//
	//mean = float32(sum)/float32(count)
	//fmt.Printf("mean 的值为: %f\n",mean)
}
