package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"log"
	"bufio"
	"time"
	"sync"
	"wiloon.com/golang-x/goroutinex"
)

var connections []net.Conn
var listener net.Listener
var (
	port    = flag.String("port", "7000", "The port to bind to")
	timeout = flag.Int("timeout", 9999, "timeout(s)")
)

func main() {
	flag.Parse()

	go startServer()

	log.Println("timeout:", *timeout)
	time.Sleep(time.Second * time.Duration(*timeout))
	for _, conn := range connections {
		conn.Close()
	}
	listener.Close()
}

func startServer() {
	flag.Parse()

	var err error
	listener, err = net.Listen("tcp4", ":" + *port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}

	defer listener.Close()
	log.Println("Listening on port:" + *port)

	for {
		var err1 error
		var conn net.Conn
		conn, err1 = listener.Accept()
		if err1 != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		connections = append(connections, conn)
		//logs an incoming message
		log.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new goroutine.
		var wg sync.WaitGroup
		wg.Add(2)
		go handleWrite(conn, &wg)
		go handleRead(conn, &wg)
	}
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer closeConn(conn, wg)
	log.Println("writer goid:", goroutinex.Goid())

	for {
		str := "hello client\r\n"
		log.Print("writing to client:" + str)

		_, e := conn.Write([]byte(str))
		if e != nil {
			log.Println("failed to send message:", e.Error())
			return
		}
		time.Sleep(time.Second * 1)
	}
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer closeConn(conn, wg)
	log.Println("reader goid:", goroutinex.Goid())

	reader := bufio.NewReader(conn)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Print("failed to read message:", err)
			return
		}
		log.Println("read msg from client:", string(line))
	}
}

func closeConn(conn net.Conn, wg *sync.WaitGroup) {
	log.Printf("gid:%v,closing conn:%v", goroutinex.Goid(), conn)
	conn.Close()

	wg.Done()
	log.Println("wg done")
}
