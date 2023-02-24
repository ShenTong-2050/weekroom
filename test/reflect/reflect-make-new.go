package main

import (
	"reflect"
	"strings"
)

func add(args []reflect.Value) (results []reflect.Value) {
	if len(args) == 0 {
		return nil
	}

	var ret reflect.Value

	switch args[0].Kind() {
	case reflect.Int:
		n := 0
		for _,a := range args {
			n += int(a.Int())
		}
		ret = reflect.ValueOf(n)
	case reflect.String:
		ss := make([]string,0,len(args))
		for _,v := range args {
			ss = append(ss,v.String())
		}
		ret = reflect.ValueOf(strings.Join(ss,""))
	}
	results = append(results,ret)
	return
}

func makeAdd(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(),add)
	fn.Set(v)		// 指向通用算法函数
}

func main() {

	var intAdd func(x,y int) int
	var strAdd func(x,y string) string

	makeAdd(&intAdd)
	makeAdd(&strAdd)

	intAdd(100,200)
	strAdd("hello","world")
}
