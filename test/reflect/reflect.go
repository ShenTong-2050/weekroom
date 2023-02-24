package main

import (
	"fmt"
	"reflect"
)

type X int

func main() {
	var num X = 100
	t := reflect.TypeOf(num)
	fmt.Println(t)
	fmt.Println(t.Name(),t.Kind())

	// 使用 reflect.ArrayOf(EleNum,ArrType) 与 reflect.MapOf(KeyType,ValType) 构造 基础 复合 类型
	a := reflect.ArrayOf(10,reflect.TypeOf(byte(0)))
	m := reflect.MapOf(reflect.TypeOf(""),reflect.TypeOf(0))
	fmt.Println(a)
	fmt.Println(m)

	// 传入对象应区分 基类型 和 指针类型
	x := 100
	tx, px := reflect.TypeOf(x),reflect.TypeOf(&x)
	fmt.Println(tx,px)
	fmt.Println(tx.Kind(),px.Kind())
	fmt.Println(tx == px.Elem())
}
