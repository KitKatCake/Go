package main

import (
	"fmt"
	"net/http"
)

func Hello2(response http.ResponseWriter,request *http.Request)  {
	fmt.Fprint(response,"Hello,Welcome to go web programming...")
}
func main()  {
	//fmt.Println("Hello Go!")

	http.HandleFunc("/",Hello2)
	http.ListenAndServe(":8080",nil)
}
