package main

import (
	"log"
	"crypto/tls"
	"net"
	"bufio"
)

func main() {
	log.SetFlags(log.Lshortfile)
	crt := "/home/roy/tmp/server.crt"
	//crt = "cert.pem"
	key := "/home/roy/tmp/server.key"
	//key = "key.pem"

	cer, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", ":8443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		println(msg)

		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
