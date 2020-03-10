package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"golang-x/goroutinex"
	"log"
	"net"
	"sync"
	"time"
)

var conn net.Conn
var (
	serverAddress = flag.String("server", "localhost:7000", "server address")
	timeout       = flag.Int("timeout", 9999, "timeout(s)")
)

func main() {

	flag.Parse()
	go startTlsClient()
	time.Sleep(time.Second * time.Duration(*timeout))
	conn.Close()
	//os.Exit(1)
}
func startTlsClient() {
	log.SetFlags(log.Lshortfile)

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	var err error
	conn, err = tls.Dial("tcp", *serverAddress, conf)
	if err != nil {
		log.Println(err)
		return
	}
	connn := conn.(net.Conn)
	log.Println("log connn:", connn)
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go handleWrite(connn, &wg)
	go handleRead(conn, &wg)
	wg.Wait()

	//n, err := conn.Write([]byte("hello\n"))
	//if err != nil {
	//	log.Println(n, err)
	//	return
	//}
	//log.Println("conn write.")
	//
	//buf := make([]byte, 100)
	//n, err = conn.Read(buf)
	//if err != nil {
	//	log.Println(n, err)
	//	return
	//}
	//
	//log.Println("conn read:", string(buf[:n]))
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer closeConn(conn, wg)

	for i := 0; i < 100; i++ {
		str := "hello server\r\n"
		_, e := conn.Write([]byte(str))
		if e != nil {
			log.Println("Error to send message because of ", e.Error())
			return
		}
		log.Print("write:" + str)
		time.Sleep(time.Second * 1)
	}
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer closeConn(conn, wg)
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
func closeConn(conn net.Conn, wg *sync.WaitGroup) {
	log.Printf("gid:%v,closing conn:%v", goroutinex.Goid(), conn)
	conn.Close()

	wg.Done()
	log.Println("wg done")
}
