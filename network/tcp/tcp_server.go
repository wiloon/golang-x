package tcp

import (
	"flag"
	"net"
	"os"
	"bufio"
	"time"
	"log"
)

func DefaultServer() {

	var address = flag.String("address", ":6789", "server address host:port")
	flag.Parse()

	listener, err := net.Listen("tcp4", *address)
	if err != nil {
		log.Println("Error listening:", err)
		os.Exit(1)
	}

	defer listener.Close()

	log.Println("Listening on " + *address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting: ", err)
			os.Exit(1)
		}
		//logs an incoming message
		log.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new go routine.
		go readFromClient(conn)
		//go writeToClient(conn)
	}
}
func readFromClient(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		log.Println("receiving...")
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			log.Print("Error to read message because of ", err)
			return
		} else {
			log.Println("receive:", string(line))
			log.Println("isPrefix:", isPrefix)
		}

	}
}
func writeToClient(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := conn.Write([]byte("hello client.\r\n"))
		if err != nil {
			log.Println("Error to send message because of ", err.Error())
			return
		}
		log.Println("sent hello to client.")
		time.Sleep(time.Second)
	}

}
