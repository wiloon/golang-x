package main

import (
	"net"
	"fmt"
	"bufio"
	"log"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		reader := bufio.NewReader(c)
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Print("failed to read message:", err)
			return
		}
		log.Println("read msg from client:", string(line))
	}
}

func test() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}
