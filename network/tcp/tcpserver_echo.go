package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)
//
//var host = flag.String("host", "", "host")
//var port = flag.String("port", "6666", "port")
//
//func main() {
//	flag.Parse()
//
//	listener, err := net.Listen("tcp", *host + ":" + *port)
//	if err != nil {
//		fmt.Println("Error listening:", err)
//		os.Exit(1)
//	}
//	defer listener.Close()
//	fmt.Println("Listening on " + *host + ":" + *port)
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			fmt.Println("Error accepting: ", err)
//			os.Exit(1)
//		}
//		//logs an incoming message
//		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
//		// Handle connections in a new go routine.
//		go handleRequest(conn)
//	}
//}
//func handleRequest(conn net.Conn) {
//	defer conn.Close()
//	for {
//		io.Copy(conn, conn)
//
//	}
//}