package main

import (
	"fmt"
	"reflect"
)

type C struct {}

func (C) Test(x,y int) (int,error) {
	return x+y,fmt.Errorf("error: %d",x+y)
}

// Format 对于 变参 来说，用 CallSlice 更方便
func (C) Format(s string, a ...interface{}) string {
	return fmt.Sprintf(s,a...)
}

func main() {
	var a C
	av := reflect.ValueOf(&a)
	// 通过 value 返回的 MethodByName 接口 传入 方法名 获取该方法
	am := av.MethodByName("Test")
	in := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(3),
	}
	out := am.Call(in)
	for _,v := range out {
		fmt.Println(v)
	}

	// 变参 示例
	var c C
	cv := reflect.ValueOf(&c)
	cm := cv.MethodByName("Format")

	// 通过 Call 调用
	out = cm.Call([]reflect.Value{
		reflect.ValueOf("%s = %d"),
		reflect.ValueOf("i"),
		reflect.ValueOf(100),
	})
	fmt.Println("Call=>",out)

	// 通过 CallSlice 调用
	out = cm.CallSlice([]reflect.Value{
		reflect.ValueOf("%s = %d"),
		reflect.ValueOf([]interface{}{"i",100}),	// 仅 一个 []interface{}{} 即可
	})
	fmt.Println("CallSlice=>",out)
}
