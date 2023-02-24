package main

import (
	"fmt"
	"sync"
	"time"
)

// Data 将读写锁嵌入到结构体
type Data struct {
	sync.RWMutex
}

type cache struct {
	data []int
	sync.RWMutex
}

// WrongMutexDemo Data 所实现的方法
func (d *Data) WrongMutexDemo(s string) {
	d.Lock()
	defer d.Unlock()
	for i:=0; i<5; i++ {
		time.Sleep(time.Second)
		fmt.Println(s,i)
	}
}

// count 重复加锁示例
func (c *cache) count() int {
	c.Lock()
	n := len(c.data)
	c.Unlock()
	return n
}

func (c *cache) get() int {
	c.Lock()
	defer c.Unlock()
	var d int
	//if n := c.count(); n > 0 {
		d = c.data[0]
		c.data = c.data[1:]
	//}
	return d
}

func main() {
	/*var d Data
	var wg sync.WaitGroup
	wg.Add(2)

	// 读协程
	go func() {
		defer wg.Done()
		d.WrongMutexDemo("read")
	}()
	// 写协程
	go func() {
		defer wg.Done()
		d.WrongMutexDemo("write")
	}()
	wg.Wait()*/

	// 重复加锁代码示例
	var c = cache{data: []int{1,2,3,4}}
	fmt.Println(c.get())
}
