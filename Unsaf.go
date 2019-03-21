package main

import (
	"fmt"
	//"reflect"
	"unsafe"
)


type user1 struct{
	b byte
	i int32
	j int64
}

type user2 struct {
	b byte
	j int64
	i int32
}

type user3 struct {
	i int32
	b byte
	j int64
}

type user4 struct {
	i int32
	j int64
	b byte
}

type user5 struct {
	j int64
	b byte
	i int32
}

type user6 struct {
	j int64
	i int32
	b byte
}
func main()  {
	//fmt.Println(unsafe.Sizeof(true))
	//fmt.Println(unsafe.Sizeof(int8(0)))
	//fmt.Println(unsafe.Sizeof(int16(10)))
	//fmt.Println(unsafe.Sizeof(int32(10000000)))
	//fmt.Println(unsafe.Sizeof(int64(10000000000000)))
	//fmt.Println(unsafe.Sizeof(int(10000000000000000)))



	//var b bool
	//var i8 int8
	//var i16 int16
	//var i64 int64
	//
	//var f32 float32
	//
	//var s string
	//
	//var m map[string]string
	//
	//var p *int32

	//fmt.Println(unsafe.Alignof(b))
	//fmt.Println(unsafe.Alignof(i8))
	//fmt.Println(unsafe.Alignof(i16))
	//fmt.Println(unsafe.Alignof(i64))
	//fmt.Println(unsafe.Alignof(f32))
	//fmt.Println(unsafe.Alignof(s))
	//fmt.Println(unsafe.Alignof(m))
	//fmt.Println(unsafe.Alignof(p))
	//fmt.Println(reflect.TypeOf(p).Align())

	fmt.Printf("------------------------------\n")

	var u1 user1
	var u2 user2
	var u3 user3
	var u4 user4
	var u5 user5
	var u6 user6

	//fmt.Println(unsafe.Offsetof(u1.b))
	//fmt.Println(unsafe.Offsetof(u1.i))
	//fmt.Println(unsafe.Offsetof(u1.j))
	fmt.Println("u2 size is ",unsafe.Sizeof(u1))
	fmt.Println("u2 size is ",unsafe.Sizeof(u2))
	fmt.Println("u3 size is ",unsafe.Sizeof(u3))
	fmt.Println("u4 size is ",unsafe.Sizeof(u4))
	fmt.Println("u5 size is ",unsafe.Sizeof(u5))
	fmt.Println("u6 size is ",unsafe.Sizeof(u6))


    i := 10
    ip := &i
    var fp *float64 = (*float64)(unsafe.Pointer(ip))

    *fp *=3

    fmt.Println(i)

    fmt.Println(&i)

    fmt.Println(ip)


}
