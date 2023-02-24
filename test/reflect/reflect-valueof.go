package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Name  string
	age   int
}

func main() {
	var a = 100
	av,ap := reflect.ValueOf(a),reflect.ValueOf(&a).Elem()
	fmt.Println(av.CanAddr(),av.CanSet())
	fmt.Println(ap.CanAddr(),ap.CanSet())

	var u = new(User)
	uv := reflect.ValueOf(u).Elem()
	name := uv.FieldByName("Name")
	age := uv.FieldByName("age")
	//fmt.Printf("the name isaddressable is:%v,iscanset is:%v\n",name.CanAddr(),name.CanSet())
	//fmt.Printf("the age isaddressable is:%v,iscanset is:%v\n",age.CanAddr(),age.CanSet())
	fmt.Println(name.CanAddr(),name.CanSet())
	fmt.Println(age.CanAddr(),age.CanSet())
	if name.CanSet() {
		name.SetString("ShenTong")
	}
	if age.CanAddr() {
		*(*int)(unsafe.Pointer(age.UnsafeAddr())) = 100
	}
	fmt.Printf("the u is: %+v\n",*u)
	//fmt.Println(u)

	fmt.Println("==========")

	var nu = User{"zhang san",18}
	nup := reflect.ValueOf(&nu)
	if !nup.CanInterface() {
		fmt.Println("nup can't used")
		return
	}
	p,ok := nup.Interface().(*User)
	if !ok {
		fmt.Println("interface failed")
		return
	}
	p.age++

	fmt.Printf("%+v\n",nu)

	fmt.Println("--------------")

	ch := make(chan int, 3)
	chv := reflect.ValueOf(ch)
	// 将 100 发送给 通道 chv 但不会 阻塞
	// 如果 chv 的 Kind 不是通道 则会 引发 panic
	// TrySend 会报告 100 是否 发送成功
	// 100 的 类型 必须是通道 所 允许的 类型
	if chv.TrySend(reflect.ValueOf(100)) {
		// 尝试 从 通道 chv 接收 一个 值 但 不会 阻塞
		// 如果 chv 的 kind 不是 chan 则会引发 panic
		// 如果 接收时 引起 阻塞 则会返回 0 false
		// 如果 通道关闭 则会返回 0 false
		x,ok := chv.TryRecv()
		fmt.Println(x,ok)
	}

	var aNil interface{} = nil
	var bNil interface{} = (*int)(nil)
	fmt.Println(aNil == nil)
	fmt.Println(bNil == nil,reflect.ValueOf(bNil).IsNil())

	s := reflect.ValueOf(struct {
		name   string
		age    int
	}{})
	//sv := reflect.ValueOf(s)
	fmt.Println(s.FieldByName("name").IsValid())
}
