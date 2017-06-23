package tcp

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

func DefaultClient() {
	var address = flag.String("address", "localhost:9000", "server address host:port")

	flag.Parse()

	conn, err := net.Dial("tcp4", *address)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("Connecting to " + *address)

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
		fmt.Println("Error to send message because of ", e.Error())
		return
	}
	fmt.Print("write:" + str)
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	for {
		fmt.Println("receiving...")
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			fmt.Print("Error to read message because of ", err)
			return
		} else {
			fmt.Println("receive:", string(line))
			fmt.Println("isPrefix:", isPrefix)
		}
	}
}
