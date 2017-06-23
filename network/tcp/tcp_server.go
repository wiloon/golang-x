package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"bufio"
	"time"
)

func TcpServer() {

	var address = flag.String("address", "localhost:9000", "server address host:port")
	flag.Parse()

	listener, err := net.Listen("tcp4", *address)
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
		go readFromClient(conn)
		//go writeToClient(conn)
	}
}
func readFromClient(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		fmt.Println("receiving...")
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			fmt.Print("Error to read message because of ", err)
		} else {
			fmt.Println("receive:", string(line))
			fmt.Println("isPrefix:", isPrefix)
		}

	}
}
func writeToClient(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := conn.Write([]byte("hello client."))
		if err != nil {
			fmt.Println("Error to send message because of ", err.Error())

		}
		fmt.Println("sent hello to client.")
		time.Sleep(time.Second)
	}

}
