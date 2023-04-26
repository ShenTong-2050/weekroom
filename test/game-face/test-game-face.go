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
