package main

import (
	"fmt"
	"sort"
	"time"
)

func sum_1(nums ...int)  {
	fmt.Print(nums," ")
	total := 0

	for _,i := range nums{
		total += i
	}
	fmt.Println(total)
}
type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength)Swap(i,j int)  {
	s[i],s[j] = s[j],s[i]
}

func (s byLength)Less(i,j int)bool  {
	return len(s[i]) < len(s[j])
}

func main()  {
	//ticker := time.NewTicker(2*time.Second)

	//for i:=0;i<10;i++{
	//	time := <- ticker.C
	//
	//	fmt.Println(time.String())
	//}

	//go func() {
	//	for t:= range ticker.C{
	//		fmt.Println("Tick at",t)
	//	}
	//}()
	//time.Sleep(4000*time.Millisecond)
	//
	//ticker.Stop()
	//
	//fmt.Println("Ticker stoppped" )


	p := fmt.Println

	//now := time.Now()
	//
	//p(now)

	p(time.Now())

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	s1 := sort.StringsAreSorted(strs)
	fmt.Println("Sorted: ", s1)

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	sum_1(1,2)
	sum_1(1,2,3)

	nums := []int{1,2,3,4}

	sum_1(nums...)










	//timer := time.NewTimer(time.Second)
	//
	//<- timer.C
	//
	//fmt.Println("Timer 1 expired")
	//
	//timer2 := time.NewTimer(time.Second)
	//
	//go func() {
	//	<-timer2.C
	//	fmt.Println("Timer 2 expired")
	//}()
	//
	//stop2 := timer2.Stop()
	//
	//if stop2{
	//	fmt.Println("Timer 2 stopped")
	//}















}
