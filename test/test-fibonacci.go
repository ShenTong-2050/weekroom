package main

import "fmt"

// 计算斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func testFibonacci(n int) int {
	const mod = 1e9 + 7
	if n<2{
		return n
	}

	p,q,r := 0,0,1

	for i :=2; i<=n; i++ {
		fmt.Println(p)
		p = q
		q = r
		r = (p + q) % mod
	}


	return r
}

func optimizeFibonacci(n int) int {
	if n < 2 {
		return n
	}

	var arr []byte

	p,q,r := 0,0,1

	for i:=2;i<=n;i++ {
		arr = append(arr,byte(i))
		p = q
		q = r
		r = p + q
	}
	fmt.Println(arr)
	fmt.Println(len(arr))
	return r
}

func main() {
	fmt.Println(fibonacci(20))
	fmt.Println(testFibonacci(20))
	fmt.Println(optimizeFibonacci(20))
}
