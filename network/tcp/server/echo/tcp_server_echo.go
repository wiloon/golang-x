package main

import (
	"net"
	"fmt"
	"os"
	"log"
	"bufio"
)

const port = "3000"

func main() {
	listener, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		fmt.Println("could not listening:", err)
		os.Exit(1)
	}

	defer listener.Close()
	log.Println("Listening on port:", port)
	conns := clientConns(listener)

	for {
		go handleConn(<-conns)
	}
}

func handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		log.Println("msg:", string(line))
		conn.Write(line)
	}

}

func clientConns(listener net.Listener) chan net.Conn {
	channel := make(chan net.Conn)
	i := 0
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("could not accept:", err)
			}
			i++
			log.Printf("%v: %v <-> %v", i, conn.LocalAddr(), conn.RemoteAddr())
			channel <- conn
		}
	}()

	return channel
}
