package main

import (
	"fmt"
)

func main() {

	var a = make([]int,0,1024)

	/*b := append(a,2)
	c := append(a,1)*/

	fmt.Println(a)
	// 在没有扩容之前 b 与 c 指向的底层切片 都是 a
	/*fmt.Println(b)
	fmt.Println(c)*/
	// 由于 d 操作给 a 中追加了三个元素 超过了 a 切片的容量, 所以拷贝了一份底层数组
	// 在没有超过 1024 个元素时 capacity 是成倍增加、如果超过 1024 则是以 cap/4 的指数增加
	e := []int{1,2}
	d := append(a,e...)
	fmt.Println(d)
	for i:=0; i<1288; i++ {
		d = append(d,i)
	}
	fmt.Println(cap(d))
}
