package main

import (
	"fmt"
	"time"
)

//func myFunc(w http.ResponseWriter,r *http.Request)  {
//	fmt.Println(w,"hi")
//}

func main()  {
	//http.HandleFunc("/", myFunc)
	//http.ListenAndServe(":8080", nil)
	var i int
	var j int
	start := time.Now()
	for i=0;i<300000000;i++{
		j += i
	}
	fmt.Println(j)
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)
}
