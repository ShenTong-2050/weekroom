package main

import (
	"errors"
	"fmt"
)

func main() {

	stack := make([]int,0,5)

	// 入栈操作
	push := func(x int) error {
		n := len(stack)
		if len(stack) == cap(stack) {
			return errors.New("the stick is full")
		}
		stack = stack[:n+1]
		stack[n] = x
		return nil
	}

	pop := func() (int,error) {
		n := len(stack)
		if n == 0 {
			return 0,errors.New("the stack is empty")
		}
		x := stack[n-1]
		stack = stack[:n-1]
		return x,nil
	}

	// 测试入栈
	for i:=0; i<7; i++ {
		fmt.Printf("push %d:%v, %v\n",i,push(i),stack)
	}

	// 测试出栈
	for i:=0; i<7; i++ {
		popX,err := pop()
		fmt.Printf("pop %d:%v %v, %v\n",i,popX,err,stack)
	}

}
