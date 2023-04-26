### 面试总结

---

#### 代码内容：

```go
package main

import "net"

func main() {
	ln,err := net.Listen("tcp",":8080")
	if err != nil {
		// handler error
	}
	for {
		conn,err := ln.Accept()
		if err != nil {
			// handler error
		}
		go handlerConnection(conn)
	}
}

func handlerConnection(conn net.Conn) {
	conn.Close()
}
```

> ###### Q：请仔细阅读如上程序，当程序编译后运行在Linux 系统上时，是否会产生用户态与内核态的切换，并说明切换时 `net.listen()` `ln.Accept()` 函数运行是用户态还是内核态

当程序在 Linux 环境 调用 `net.Listen()` 与 `ln.Accept()` 函数时会触发系统调用 由 用户态 切换到 内核态。内核负责创建所有网络请求并返回描述符，同时在运行期间处理所有网络 IO。

当调用 `net.Listen()` 函数时涉及 socket 的创建 与 端口的绑定 会发生 用户态 与 内核态 的切换。

调用 `ln.Accept()` 函数时会发生阻塞，等待客户端请求的到来，在客户端连接成功之后 内核 将返回 已连接套接字 的 描述符，从而实现 内核态 到 用户态 的切换。

在执行 `go handlerConnection()`函数期间由于所有的网络 IO 都已经有内核处理完毕，所以 `handlerConnetion()` 函数的运行处于 用户态。

