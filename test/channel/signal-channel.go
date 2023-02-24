package main

import (
	"os"
	"os/signal"
	"sync"
)

var exits = &struct {
	sync.RWMutex
	signals   chan os.Signal
	funcS    []func()
}{}

func atExits(f func()) {
	exits.Lock()
	defer exits.Unlock()
	exits.funcS = append(exits.funcS,f)
}

func waitExits() {
	if exits.signals == nil {
		exits.signals = make(chan os.Signal)
		//signal.Notify(exits.signals,syscall.SIGINT,syscall.SIGTERM)
		signal.Notify(exits.signals)
	}
	exits.RLock()
	for _,f := range exits.funcS {
		defer f()		// defer 按照 FILO 顺序执行
	}					// 即使某些函数panic,defer 也能确保后续函数正常执行
	exits.RUnlock()
	<-exits.signals
}

func main() {
	atExits(func() {println("exit1...")})
	atExits(func() {println("exit2...")})
	atExits(func() {println("exit3...")})
	waitExits()
}
