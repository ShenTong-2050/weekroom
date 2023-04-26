package main

import (
	"fmt"
	"time"
)

func sum(a,b int) (result int) {
	defer trace("sum")()
	// defer func() { result+=a }()
	result+=a
	result+=b
	time.Sleep(time.Second * 2)
	return result
}

func trace(msg string) func() {
	start := time.Now()
	fmt.Printf("start %s time %v\n",msg,start.Unix())
	return func() {
		fmt.Printf("end %s time %v, cost time %v\n",msg,time.Now().Unix(),time.Since(start))
	}
}

func main() {
	count := sum(100,200)
	fmt.Printf("%d\n",count)
}
