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

> ##### Q：请仔细阅读如上程序，当程序编译后运行在Linux 系统上时，是否会产生用户态与内核态的切换，并说明切换时 `net.listen()` `ln.Accept()` 函数运行是用户态还是内核态

当程序在 Linux 环境 调用 `net.Listen()` 与 `ln.Accept()` 函数时会触发系统调用 由 用户态 切换到 内核态。内核负责创建所有网络请求并返回描述符，同时在运行期间处理所有网络 IO。

当调用 `net.Listen()` 函数时涉及 socket 的创建 与 端口的绑定 会发生 用户态 与 内核态 的切换。

调用 `ln.Accept()` 函数时会发生阻塞，等待客户端请求的到来，在客户端连接成功之后 内核 将返回 已连接套接字 的 描述符，从而实现 内核态 到 用户态 的切换。

在执行 `go handlerConnection()`函数期间由于所有的网络 IO 都已经有内核处理完毕，所以 `handlerConnetion()` 函数的运行处于 用户态。

> #### Q：当 `net.listen()` 运行后本进程的监听文件存储在什么位置?以及监听文件的信息有哪些?

当 `net.Listen()` 运行之后本进程的监听文件存储在 `/proc/net/tcp` 或 `/proc/net/tcp6（IPV6）` 文件中。
这些文件是 Linux 内核用于记录网络连接的文件之一。

监听文件的信息包括：

* inode（监听文件的唯一标识符）
* 本地 IP 地址 和 端口号
* 远程 IP 地址 和 端口号
* TCP 连接状态【如 TCP_ESTABLISHED】
* Socket 选项【如 SO_REUSEADDR等】

> #### Q：当for运行接收到100个网络连接以后，是否会生成新文件?生成多少个文件?它们的连接信息是怎么表示的?

在for循环接收到100个网络连接后，不会生成新的文件。每个网络连接是通过一个套接字（socket）来表示的，
套接字是一种通信机制，它是通过一个文件描述符（file descriptor）来引用的。

在Linux系统中，所有的设备、文件、网络套接字都是用文件描述符来表示的，因此这些套接字信息也是存储在文件描述符中的。

在这段代码中，每次循环生成的conn变量就是一个网络连接的套接字，它所包含的信息包括源IP地址、源端口号、目的IP地址、目的端口号等。

> #### Q：上面的go handlerConnection(conn) 启动新子协程以后，请说明GO GMP调度过程

当 go handlerConnection(conn) 这段代码被执行时，它会创建一个新的goroutine来执行 handlerConnection 函数。

在Go语言中，goroutine的调度是由Go的调度器（GMP调度器）来完成的。

GMP调度器中的G代表Goroutine，M代表Machine，P代表Processor。
当有新的goroutine被创建时，调度器会将其分配到其中一个M中，然后该M会被分配到其中一个P中，并在P的运行队列中排队。
当一个goroutine被调度执行时，该P的运行队列会出列一个goroutine，并将其交给该P上的M来执行。
如果一个goroutine阻塞了，那么它会被从M中移除，并被重新放入到等待队列中。
调度器会选择另一个goroutine来执行，直到该goroutine被唤醒。

在这个例子中，当 go handlerConnection(conn) 被执行时，它会创建一个新的goroutine，然后该goroutine会被分配到一个M中，
并放入到运行队列中等待执行。当P的运行队列中没有其他goroutine时，该goroutine会被立即执行。
当该goroutine被阻塞时，调度器会将其从M中移除，并放入到等待队列中，直到有新的goroutine可供执行。

> #### 请以G0伪代码写出如何实现个 负载均衡转 发请求和响应的游戏网关API服务器

```go
// 负载均衡器服务器
func LoadBalancerServer() {
    // 连接游戏网关API服务器的地址列表
    servers := []string{"api-server-1.com", "api-server-2.com", "api-server-3.com"}

    // 监听负载均衡器服务器端口
    listener, _ := net.Listen("tcp", ":8080")

    for {
        // 接受客户端请求
        clientConn, _ := listener.Accept()

        // 从游戏网关API服务器列表中选择一个服务器
        server := servers[rand.Intn(len(servers))]

        // 转发请求到游戏网关API服务器
        go ForwardRequest(clientConn, server)
    }
}

// 转发请求到游戏网关API服务器
func ForwardRequest(clientConn net.Conn, server string) {
    // 连接游戏网关API服务器
    serverConn, _ := net.Dial("tcp", server)

    // 复制客户端请求数据到游戏网关API服务器
    go io.Copy(serverConn, clientConn)

    // 复制游戏网关API服务器响应数据到客户端
    go io.Copy(clientConn, serverConn)
}
```

上面的示例中，LoadBalancerServer() 函数是负载均衡器服务器的主函数，它监听本地的 8080 端口，
接受客户端的请求，并随机选择一个游戏网关API服务器地址， 将请求转发到对应的服务器上，
使用 ForwardRequest() 函数来实现请求转发。
ForwardRequest() 函数负责连接游戏网关API服务器，将客户端请求数据复制到游戏网关API服务器上，
并将游戏网关API服务器响应数据复制到客户端上。
这样，客户端就可以通过负载均衡器服务器来访问游戏网关API服务器，实现了负载均衡的功能。