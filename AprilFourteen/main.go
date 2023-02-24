package main

import (
	"fmt"
)

const (
	zero = "A"
	one
	two
	three = iota
	four
	_
	six
	seven = iota + 4
)

func main() {
	var (
		name = "Test"
		age  = 18
		sex  = "man"
	)
	fmt.Println(name,age,sex)
	fmt.Println(zero,one,two,three,four,six,seven)
}
