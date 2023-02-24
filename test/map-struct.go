package main

import "fmt"

type Profile struct {
	Name 	string
	Age 	int
}

func main() {
	m := map[int]Profile{
		1:{"test",18},
	}
	fmt.Println(m)
	v := Profile{Name: "edit_test",Age: 20}	// 通过设置整个 value 来修改目标对象
	m[1] = v
	fmt.Println(m)

	m1 := map[int]*Profile{
		1:{"ptr profile",19},
	}
	m1[1].Age = 22			// m1[1] 返回的是指针,可通过指针修改目标对象
	fmt.Println(m1[1])
}
