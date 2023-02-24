package main

import "fmt"

type N int

func (n N) Value() {
	n++
	fmt.Printf("ptr : %p, value:%v\n", &n, n)
}

func (n *N) Pointer() {
	*n++
	fmt.Printf("ptr : %p, value:%v\n", n, *n)
}

func main() {

	var a N = 11

	p := &a

	a.Value()
	a.Pointer()

	p.Value()
	p.Pointer()

	fmt.Printf("ptr : %p, value:%v\n", &a, a)
}
