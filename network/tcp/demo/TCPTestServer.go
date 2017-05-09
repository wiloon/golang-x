package demo

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "7000", "port")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", *host + ":" + *port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Listening on " + *host + ":" + *port)
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
	index := 0;
	for {
		index = index + 1
		fmt.Println("loop:", index)
		//io.Copy(conn, conn)
		_, err := conn.Write([]byte("xxx xxx"))
		if err != nil {
			fmt.Println("Error to send message because of ", err.Error())
			break
		}
		time.Sleep(time.Second)
	}
}