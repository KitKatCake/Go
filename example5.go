package main

import (
	"sync"
	"fmt"
	"time"
)

var once sync.Once

type User struct{
	name string
	age int
}

type Book struct{
	BookName string
	L *sync.Mutex

}

func (bk *Book) SetName(wg *sync.WaitGroup,name string)  {
	defer func() {
		fmt.Println("Unlock set name:", name)
		bk.L.Unlock()
		wg.Done()
	}()

	bk.L.Lock()

	fmt.Println("Lock set name:",name)
	time.Sleep(time.Second)
	bk.BookName = name

}

type A struct{
	Face int
}

type B interface {
	f()
}

func (a A) f() {
	fmt.Println("hi ",a.Face)
}

//func myfunc(w http.ResponseWriter,r *http.Request)  {
//	fmt.Println(w,"hi")
//}

func onces()  {
	fmt.Println("onces")
}

func onced()  {
	fmt.Println("onced")
}

func main()  {
	//p := User{"John",28}
	//fmt.Printf("Printf struct %%v : %v\n",  p)

	//http.HandleFunc("/",myfunc)
	//http.ListenAndServe(":8080",nil)

	//var s A = A{9}

	//s.f()

	//var b B = A{9}

	//b.f()




	//bk := Book{}
	//
	//bk.L = new(sync.Mutex)
	//
	//wg := &sync.WaitGroup{}
	//
	//books := []string{"《三国演义》", "《道德经》", "《西游记》"}
	//
	//for _,book := range books{
	//	wg.Add(1)
	//	go bk.SetName(wg, book)
	//}
	//wg.Wait()



	for i,v := range make([]string,10){
		once.Do(onces)
		fmt.Println("v:",v,"---i:",i)
	}

	for i := 0;i<10;i++{
		go func(i int) {
			once.Do(onced)
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second*4)

}
