package main

import (
	"sync"
	"testing"
)

var m sync.Mutex

func call() {
	m.Lock()
	m.Unlock()
}

func deferCall() {
	m.Lock()
	// 延迟执行导致耗时变长
	defer m.Unlock()
}

func BenchmarkCall(t *testing.B) {
	for i:=0; i<t.N; i++ {
		call()
	}
}

func BenchmarkDefer(t *testing.B) {
	for i := 0; i<t.N; i++ {
		deferCall()
	}
}



