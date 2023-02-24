package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)

	c := make(chan int)

	go func() {
		defer wg.Done()
		for {
			num,ok := <-c
			if !ok {
				return
			}
			fmt.Println(num)
		}
	}()

	for i := 0; i<10; i++ {
		c <-i
	}

	close(c)

	wg.Wait()

	fmt.Println("done")

}

