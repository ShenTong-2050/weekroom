package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var Wg sync.WaitGroup

func count() {
	var x = 0
	for i:=0; i<math.MaxInt32; i++ {
		x += i
	}
	println(x)
}

func SingleCpuCount(n int) {
	for i:=0; i<n; i++ {
		count()
	}
}

func GoroutineCpuCount(n int)  {
	for i:=0; i<n; i++ {
		go func() {
			count()
			Wg.Done()
		}()
	}
	Wg.Wait()
}

func main() {
	n := runtime.GOMAXPROCS(0)
	fmt.Println("default Gomaxprocs is:",n)
	//SingleCpuCount(n)
	GoroutineCpuCount(n)
}
