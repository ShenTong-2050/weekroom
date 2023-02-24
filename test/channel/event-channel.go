package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

func main() {

	/*// 事件通道
	done := make(chan struct{})

	// 数据传输通道
	strChn := make(chan string)

	go func() {
		s := <- strChn	// 接收数据
		fmt.Println(s)
		close(done)
	}()

	// 给通道发送数据
	strChn<-"hello world"

	// 通道阻塞直到有数据或通道关闭
	<-done*/

	//noClogChan()
	//bufferChan()
	//isSync()
	//okChan()
	//publicInform()
	singleChan()
}

// 非阻塞通道
func noClogChan() {

	ch := make(chan int,3)

	ch<-1
	ch<-2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

// 通道缓冲区对象
func bufferChan() {
	var a,b chan int = make(chan int,3),make(chan int)
	var c chan bool
	println(a==b)
	println(c==nil)
	fmt.Printf("%p,%d\n",a,unsafe.Sizeof(a))
}

// 通过缓冲区大小和已缓冲数量来判断是同步还是异步
func isSync() {
	a,b := make(chan int),make(chan int,3)

	b<-1
	b<-2

	println("a:",len(a),cap(a))
	println("b:",len(b),cap(b))
}

func okChan() {
	done := make(chan struct{})
	data := make(chan int)
	go func() {
		defer close(done)		// 确保发出结果通知
		// 接收方式一
		/*for {
			i,ok := <-data
			if !ok {			// 据此判断通道是否关闭
				return
			}
			fmt.Println(i)
		}*/
		// 接收方式二
		for x := range data {
			fmt.Println(x)
		}
	}()
	data<-1
	data<-2
	data<-3
	close(data)
	<-done
}

// 群发性通知
func publicInform() {
	var wg sync.WaitGroup

	var ready = make(chan struct{})

	for i:=0; i<3; i++ {
		// 计数器加 1
		wg.Add(1)

		go func(id int) {
			// 计数器减一
			defer wg.Done()
			fmt.Println(id,"ready...")
			<-ready
			fmt.Println(id,"running...")
		}(i)
	}
	// 睡一秒给 goroutine 运行时间
	time.Sleep(time.Second * 2)
	println("Ready????Go...")
	// 关闭通道、阻塞关闭、继续运行 goroutine
	close(ready)
	// 等待计数器归零
	wg.Wait()
}

// 单项通道
func singleChan() {
	// 初始化一个计数队列
	var wg sync.WaitGroup
	// 增加两个基数操作【发送与接收】
	wg.Add(2)
	// 创建一个双向通道
	var ch = make(chan int)
	// 通过类型转换将双向通道的 发送 操作赋值给 send
	var send chan <- int = ch
	// 通过类型转换将双向通道的 接收 操作赋值给 receive
	var receive <- chan int = ch

	// 接收操作
	go func(){
		// 计数器减 1
		defer wg.Done()
		// 接收通道数据
		for val := range receive{
			fmt.Println(val)
		}
	}()

	// 发送操作
	go func() {
		// 计数器减 1
		defer wg.Done()
		// 关闭通道 ch
		defer close(ch)
		// 给 send 通道发送 3 个数字
		for i:=0; i<3; i++ {
			send <- i
		}
	}()

	// 等待计数器归零
	wg.Wait()
}
