package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"sync"
	"time"
	"wiloon.com/golang-x/goroutinex"
)

var conn net.Conn

var (
	serverAddress = flag.String("server", "localhost:3000", "server address")
	timeout       = flag.Int("timeout", 9999, "timeout(s)")
)

func main() {
	flag.Parse()
	go startClient()
	log.Println("timeout:", *timeout)
	time.Sleep(time.Second * time.Duration(*timeout))
	conn.Close()
	//os.Exit(1)
}

func startClient() {
	flag.Parse()
	var err error
	conn, err = net.Dial("tcp4", *serverAddress)
	if err != nil {
		log.Println("Error connecting:", err)
		os.Exit(1)
	}

	defer conn.Close()

	log.Println("Connecting to " + *serverAddress)

	var wg sync.WaitGroup
	wg.Add(2)
	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)
	wg.Wait()
}
func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer closeConn(conn, wg)

	for {
		str := "hello server\r\n"
		log.Print("writing to server:" + str)
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
	reader := bufio.NewReader(conn)
	for {
		log.Println("receiving from server...")
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Print("failed to read message:", err)
			return
		}
		log.Println("read msg from server:", string(line))
	}
}

func closeConn(conn net.Conn, wg *sync.WaitGroup) {
	log.Printf("gid:%v,closing conn:%v", goroutinex.Goid(), conn)
	conn.Close()

	wg.Done()
	log.Println("wg done")
}
