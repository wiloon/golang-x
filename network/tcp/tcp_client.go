package tcp

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
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

	//var str = "hello " + strconv.Itoa(i) + "\r\n"
	var str = "hello server."

	_, err := conn.Write([]byte(str))
	fmt.Println("sent: " + str)

	if err != nil {
		fmt.Println("Error to send message because of ", err.Error())
	}

}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	for {
		fmt.Println("receiving...")
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			fmt.Print("Error to read message because of ", err)
			time.Sleep(time.Second)
		} else {
			fmt.Println("receive:", string(line))
			fmt.Println("isPrefix:", isPrefix)
		}

	}
}
