package main

import (
	"fmt"
	"sync"
	"time"
)

type user_1 struct{
	lock sync.Mutex
	name string
	age int
}

type stu struct {

}

func (st *stu)speak()  {
	fmt.Println("I am  a student.")
}

type Myerror struct {
	err error
}

func (e Myerror) Error() string  {
	return e.err.Error()
}

type Istu interface {
	speak()
}

func server(ch <-chan int) {
	for val := range ch {
		go handle(val)
	}
}

func handle(x int)  {
	defer func() {
		if err := recover();err!=nil{
			fmt.Println("work failed:",err)
		}
	}()

	if x == 3 {
		panic(x)
	}

	fmt.Println("work succeeded:",x)
}

func main()  {

	//err3 := fmt.Errorf("hello error3")
	//if err3 != nil {
	//	fmt.Println(err3)
	//}
	//
	//var s3 *stu = nil
	//
	//fmt.Println("s3 == nil", s3 == nil)
	//
	//fmt.Printf("%T\n", s3)
	//
	//s2 := new(interface{})
	//
	//fmt.Println("s2 == nil", s2 == nil)
	//fmt.Println("*s2 == nil", *s2 == nil)
	//fmt.Printf("%T\n", s2)
	//fmt.Printf("%T\n", *s2)
	//
	//err2 := Myerror{
	//	errors.New("hello error2"),
	//}
	//
	//
	//fmt.Println(err2.Error())



	//var user = os.Getenv("USER")
	//
	//if user == ""{
	//	panic("The USER environment variable is not set.")
	//}


	ch := make(chan int,6)

	for i :=1;i<=6;i++{
		ch <- i
	}

	close(ch)

	go server(ch)

	time.Sleep(time.Millisecond*100)







	//	var i *int
//	i = new(int)
//	*i = 10
//
//	fmt.Println(*i)
//
//	u := new(user_1)
//
//	u.lock.Lock()
//
//	u.name = "张三"
//
//	u.age = 10
//
//	u.lock.Unlock()
//
////	fmt.Println(u)
//    error := errors.New("hello,error")
//
//    if error != nil{
//    	fmt.Println(error)
//	}

}
