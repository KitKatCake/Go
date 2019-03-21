package main

import (
	"fmt"
	"os"
)

func main()  {
	//switch  time.Now().Weekday() {
	//case time.Saturday,time.Sunday:
	//	fmt.Println("It's the weekend")
	//default:
	//	fmt.Println("It's a weekday")
	//}
	//
	//t := time.Now()
	//
	//switch  {
	//case t.Hour()<12:
	//	fmt.Println("It's before noon")
	//default:
	//	fmt.Println("It's after noon")
	//
	//}
	//
	//whatAmI := func(i interface{}){
	//	switch t:=i.(type) {
	//	case bool:
	//		fmt.Println("I'm a bool")
	//	case int:
	//		fmt.Println("I'm an int")
	//	default:
	//		fmt.Printf("Don't know type %T\n", t)
	//	}
	//}
	//
	//whatAmI(true)
	//whatAmI(1)
	//whatAmI("hey")





	f := createFile("defer.txt")
	defer closeFile(f)
	writeFile(f)




	//fmt.Println(os.Getpid())
	//fmt.Println(os.Getppid())
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
func createFile(p string) *os.File {

	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
