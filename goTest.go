package main

import "fmt"

//import "encoding/json"


type user struct{
	name string
	email string
}

type admin struct {
	user
	level string
}

func (a admin) sayHello2() {
	fmt.Println("Hello，i am a admin")
}

type Hello interface {
	hello()
}

func (u user)hello()  {
	fmt.Println("Hello，i am a user")
}
func sayHello(h Hello)  {
	h.hello()
}

func (u user) sayHello2() {
	fmt.Println("Hello，i am a user")
}
func main()  {


	//b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	//var f interface{}
	//
	//err := json.Unmarshal(b,&f)
	//
	//m := f.(map[string]interface{})


	fmt.Println("Hello")


	ad := admin{user{"张三","zhangsan@flysnow.org"},"管理员"}

	fmt.Println("可以直接调用,名字为：",ad.name)
	fmt.Println("也可以通过内部类型调用,名字为：",ad.user.name)
	fmt.Println("但是新增加的属性只能直接调用，级别为：",ad.level)

	ad.sayHello2()
	ad.user.sayHello2()

	sayHello(ad.user)

	sayHello(ad)

}
