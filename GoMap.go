package main

import (
	"fmt"
	"sync"
)

func main()  {
	var m sync.Map

	m.Store("name","Joe")
	m.Store("gender","Male")

	v,ok := m.LoadOrStore("name1","Jim")

	fmt.Println(ok,v)

	v, ok = m.LoadOrStore("name", "aaa")
	fmt.Println(ok, v)

	v,ok = m.Load("name")
	if ok{
		fmt.Println("key存在,值是: ",v)
	}else{
		fmt.Println("key不存在")
	}


	f := func(k,v interface{}) bool{
		fmt.Println(k,v)
		return true
	}

	m.Range(f)

	//m.Delete("name1")

	//fmt.Println(m.Load("name1"))






















}
