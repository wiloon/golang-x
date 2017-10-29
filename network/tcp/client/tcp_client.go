package main

import (
	"bufio"
	"flag"
	"net"
	"os"
	"sync"
	"log"
)

const address = "localhost:9000"

func main() {

	flag.Parse()

	conn, err := net.Dial("tcp4", address)
	if err != nil {
		log.Println("Error connecting:", err)
		os.Exit(1)
	}

	defer conn.Close()

	log.Println("Connecting to " + address)

	var wg sync.WaitGroup
	wg.Add(2)
	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)
	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	str := "hello server\r\n"
	_, e := conn.Write([]byte(str))
	if e != nil {
		log.Println("Error to send message because of ", e.Error())
		return
	}
	log.Print("write:" + str)
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
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
