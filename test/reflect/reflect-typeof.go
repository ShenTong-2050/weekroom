package main

import (
	"fmt"
	"reflect"
)


type user struct {
	name string `json:"name,omitempty"`
	age  int    `json:"age,omitempty"`
}

type profile struct {
	user  `json:"user"`
	title string `json:"title,omitempty"`
}

type M int

func (M) String() string {
	return ""
}


func main() {
	var p profile
	tp := reflect.TypeOf(&p)
	// 如果是指针类型
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	for i:=0; i<tp.NumField(); i++ {
		f := tp.Field(i)
		fmt.Println(f.Name,f.Type,f.Offset)
		// 如果是匿名字段
		if f.Anonymous {
			for j := 0; j<f.Type.NumField(); j++ {
				af := f.Type.Field(j)
				fmt.Println(" ",af.Name,af.Type)
			}
		}
	}

	// 通过 FieldByName 按名称查找
	name,_ := tp.FieldByName("name")
	fmt.Println(name.Name,name.Type,name.Offset)
	// 按 多级 索引 查找
	title := tp.FieldByIndex([]int{0})
	fmt.Println(title.Name,title.Type,title.Offset)

	// 获取所有字段的 tag 标记
	fmt.Println("tag...")
	for m := 0; m <tp.NumField(); m++ {
		fi := tp.Field(m)
		fmt.Println(fi.Tag.Get("json"),fi.Tag.Get("omitempty"))
		// 获取匿名字段的 tag 信息
		if fi.Anonymous {
			for n := 0; n<fi.Type.NumField(); n++ {
				fin := fi.Type.Field(n)
				fmt.Println(fin.Name,fin.Tag.Get("json"),fi.Tag.Get("omitempty"))
			}
		}
	}

	// 辅助判断方法 Implements、ConvertibleTo、AssignableTo
	var x M
	t := reflect.TypeOf(x)

	// Implements 不能直接使用类型作为参数
	st := reflect.TypeOf( (*fmt.Stringer) (nil)).Elem()
	fmt.Println(st)
	// 判断 该类型(t) 是否实现了接口 st
	fmt.Println(t.Implements(st))

	it := reflect.TypeOf(0)
	// 判断 该类型(t) 的值是否可转换为 it 类型
	fmt.Println(t.ConvertibleTo(it))

	// 判断 该类型的值(t) 是否可赋值给类型it
	fmt.Println(t.AssignableTo(st),t.AssignableTo(it))

}
