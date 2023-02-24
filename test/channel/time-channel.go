package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	go func() {
		for {
			select {
			case <-time.After(time.Second * 5):
				fmt.Println("timeout...")
				os.Exit(0)
			}
		}
	}()

	go func() {
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				fmt.Println(time.Now())
			}
		}
	}()

	// 直接用 nil channel 阻塞进程
	<-(chan struct{})(nil)

}
