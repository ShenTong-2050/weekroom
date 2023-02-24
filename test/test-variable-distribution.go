package main

func test(x *int) {

	// 指针传递导致引用周期变长并在堆上分配内存
	go func() {
		println(x)
	}()

}

// 变量分配分析
func main() {

	var i int = 100

	// 拷贝变量 i 的指针地址
	p := &i

	test(p)
}
