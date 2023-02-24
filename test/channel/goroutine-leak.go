package main

import (
	"runtime"
	"time"
)

// Leak 起10个协程读取 channel 数据
func Leak() {
	c := make(chan int)
	for i:=0; i<10; i++ {
		go func() {
			<-c
		}()
	}
}

func main() {

	Leak()

	for {
		time.Sleep(time.Second)
		// 强制垃圾回收
		runtime.GC()
	}
}
