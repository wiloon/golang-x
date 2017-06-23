package tcp

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

func EchoServer() {
	var address = flag.String("address", "localhost:9000", "server address host:port")

	flag.Parse()
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *address)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + *address)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		//logs an incoming message
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		io.Copy(conn, conn)
	}
}
