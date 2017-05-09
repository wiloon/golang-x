package demo

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"strconv"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "7777", "port")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *host + ":" + *port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + *host + ":" + *port)
	var wg sync.WaitGroup
	wg.Add(2)
	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)
	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i > 0; i-- {
		//var str = "hello " + strconv.Itoa(i) + "\r\n"
		var str = "hello " + strconv.Itoa(i)
		_, err := conn.Write([]byte(str))
		fmt.Println("sent: " + str)

		if err != nil {
			fmt.Println("Error to send message because of ", err.Error())
			break
		}
	}
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	for {
		fmt.Println("receive.")
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			fmt.Print("Error to read message because of ", err)
			return
		}
		fmt.Println("receive:", string(line))
		fmt.Println("isPrefix:", isPrefix)
	}
}