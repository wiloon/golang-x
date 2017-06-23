package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

func TcpServerEcho() {

	var address = flag.String("address", "localhost:9000", "server address host:port")

	flag.Parse()

	listener, err := net.Listen("tcp", *address)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Listening on " + *address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		//logs an incoming message
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new go routine.
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		io.Copy(conn, conn)
	}
}