package main

func consumer(data chan int, done chan bool) {
	for v := range data {	// 接收数据直到通道关闭
		println(v)
	}
	done <- true	// 通知 main、消费结束
}

func producer(data chan int) {
	for x := 0; x<5; x++ {
		data <-x	// 发送数据
	}
	close(data)		// 生产结束、关闭通道
}

func main() {

	done := make(chan bool)	// 用于接收消费者结束信号

	data := make(chan int)	// 数据管道

	go producer(data)	// 启动生产者

	go consumer(data,done)	// 启动消费者

	<-done		// 阻塞直到消费者发回结束信号
}
