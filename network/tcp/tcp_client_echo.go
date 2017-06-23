package tcp

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

func EchoClient() {

	var host = flag.String("host", "localhost", "host")
	var port = flag.String("port", "9000", "port")

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
	go handleEchoWrite(conn, &wg)
	go handleEchoRead(conn, &wg)
	wg.Wait()
}
func handleEchoWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i > 0; i-- {
		str := "hello " + strconv.Itoa(i) + "\r\n"
		_, e := conn.Write([]byte(str))
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}
		fmt.Println("write:" + str)
	}
}
func handleEchoRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	for i := 1; i <= 10; i++ {
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			fmt.Print("Error to read message because of ", err)
			return
		}
		fmt.Print("read:" + line)
	}
}
