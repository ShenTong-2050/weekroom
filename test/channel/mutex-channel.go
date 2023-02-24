package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	sync.RWMutex
	profile
}

type profile struct {
	name  string
	age   int
}

// Mutex 通过 channel 实现一个互斥锁
type Mutex struct {
	ch chan struct{}
}

// NewMutex 实现互斥锁需要先初始化
func NewMutex() *Mutex {
	mu := &Mutex{make(chan struct{},1)}
	mu.ch <- struct{}{}
	return mu
}

// Lock 加锁
func (m *Mutex) Lock(goName string) {
	// 一直读取数据
	println(goName)
	<-m.ch
}

// Unlock 解锁
func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		println("unlock of unlocked mutex")
	}
}

// TryLock 尝试获取锁
func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

// LockTimeOut 加入超时设置
func (m *Mutex) LockTimeOut(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	// 未超时
	case <-m.ch:
		timer.Stop()
		return true
	// 超时 case
	case <-timer.C:
	}
	return false
}


func main() {

	/*var d data

	go func() {
		for i:=0; i<5;i++ {
			d.Lock()
			d.profile.name = fmt.Sprintf("%s%d","zhangSan",i)
			fmt.Println(d.profile.name)
			time.Sleep(time.Second)
			d.Unlock()
		}
	}()

	go func() {
		for i:=0; i<5;i++ {
			d.Lock()
			d.profile.name = fmt.Sprintf("%s%d","liSi",i)
			fmt.Println(d.profile.name)
			time.Sleep(time.Second)
			d.Unlock()
		}
	}()*/

	// 初始化一个 容量为 1 的 chan
	mu := NewMutex()

	// 尝试获取锁
	ok := mu.TryLock()
	fmt.Printf("locked v %v\n",ok)

	mu.Unlock()
	ok = mu.TryLock()
	fmt.Printf("locked v %v\n",ok)
	mu.Unlock()

	ok1 := mu.LockTimeOut(time.Microsecond * 100)
	fmt.Printf("%v\n",ok1)

}
