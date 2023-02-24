package main

import (
	"fmt"
	"reflect"
)

type U struct {
	S
}

type S struct {}

type X int

func (U) ValU() {
	fmt.Println("this is uValue")
}
func (*U) PtrU() {
	fmt.Println("this is uPtr")
}
func (S) ValS() {
	fmt.Println("this is eVal")
}
func (*S) PtrS() {
	fmt.Println("this is ePtr")
}
func (x X) Pre() {
	fmt.Printf("pre.u : %p, value : %v\n",&x,x)
}

func (x *X) Pres() {
	fmt.Printf("pres.x : %p, value : %v\n",x,*x)
}

func methodSet(f interface{}) {
	// 返回接口中保存值的类型
	t := reflect.TypeOf(f)
	fmt.Println(t.NumMethod())
	// 遍历获取所有的方法集
	for i,n:=0,t.NumMethod(); i<n;i++ {
		m := t.Method(i)
		fmt.Println(m)
	}
}

func main() {

	//var f U
	//methodSet(f)	// 显示 user 方法集
	//fmt.Println("------------")
	//methodSet(&f)	// 显示 *user 方法集

	// 通过类型引用方法退化成函数执行，参数为显示传参，
	// 至于方法只要是实现了该类型并且在方法集内
	var x X = 10
	/*fmt.Printf("main.x : %p, value : %v\n",&x,x)
	f1 := X.Pre
	f1(x)
	f2 := (*X).Pre
	f2(&x)*/

	p := &x

	x++
	f1 := x.Pres		// 由于 Pres 的 receiver 为指针类型,所以复制指针 &x

	x++
	f2 := p.Pres		// 复制 p 指针

	x++
	fmt.Printf("main.x %p, value : %v\n",&x,x)

	f1()

	f2()
}
