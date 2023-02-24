package main

import "sync"

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	// 声明两个 channel 通道
	var a = make(chan int)

	// 接收端
	go func() {
		defer wg.Done()
		for {
			var (
				x		int			// 从通道中接收的 int 值
				ok		bool		// 判断通道是否关闭
			)
			select {				// 随机选择接收通道
			case x,ok = <- a:
				println("a1",x)
				break
			case x,ok = <- a:
				println("a2",x)
				break
			}

			if !ok {
				return
			}
		}
	}()

	// 发送端
	go func() {
		defer wg.Done()
		defer close(a)
		//defer close(b)
		for i:=0; i<10; i++ {	// 随机选择发送通道
			select {
			case a <- i:
			//case b <- i*100:
			}
		}
	}()

	wg.Wait()
}
