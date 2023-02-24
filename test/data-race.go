package main

import (
	"sync"
	"time"
)

func main() {

	// 通过读写锁实现同步,避免读写操作同时进行
	m := make(map[string]int)

	var lock = sync.RWMutex{}

	go func() {
		for {
			lock.Lock()
			m["a"] = 1
			lock.Unlock()
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			lock.Lock()
			_ = m["b"]
			lock.Unlock()
			time.Sleep(time.Microsecond)
		}
	}()

	select{}
}

func dataRace() {
	// 并发读写引起数据竞争
	m := make(map[string]int)

	// 并发写操作
	go func() {
		for{
			m["a"] = 1
			time.Sleep(time.Microsecond)
		}
	}()

	// 并发读操作
	go func() {
		for{
			_ = m["b"]
			time.Sleep(time.Microsecond)
		}
	}()

	select {}
}
