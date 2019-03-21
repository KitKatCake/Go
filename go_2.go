package main

import (
	"sync"
)
func gen2(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
func sq2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
func merge(cs ...<- chan int) <- chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output:= func(c <- chan int) {
		for n:= range c{
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))

	for _,c:=range cs{
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main()  {

	//in := gen2(2,3)
	//c1 := sq2(in)

	//c2 := sq2(in)

	//for n:=range merge(c1,c2){
	//	fmt.Println(n)
	//}



	//m := make(map[int]string)
	//
	//m[0] = "a"
	//m[1] = "b"
	//m[2] = "c"
	//
	//for _,v := range m{
	//	fmt.Println(v)
	//}
	//
	//for k,_ := range m{
	//	fmt.Println(k)
	//}

}
