package main

import (
	"fmt"
	"reflect"
)

type profile struct {
	Name 	string	`姓名`
	Sex 	int 	`性别`
	Age 	int 	`年龄`
	Address string 	`地址`
}

func (p *profile) pre() {
	fmt.Printf("%p , %v\n", p, p)
}

func main() {

	p := profile{Name: "test",Sex: 0,Age: 18,Address: "山西省"}

	v := reflect.ValueOf(p).Type()

	// 获取所有字段的标签
	for i,n := 0,v.NumField(); i<n; i++ {
		fmt.Printf("%s:%v\n",v.Field(i).Tag,v.Field(i))
	}

	s := (*profile).pre
	s(&p)

}

