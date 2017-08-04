package rpcx

// first we create a simple golang rpc server based on socket
import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Counter struct {
	Sum int
}

func (this *Counter) Add(i int, r *int) error {
	this.Sum += i
	*r = this.Sum
	fmt.Printf("i: %v", i)
	return nil
}

func NewJsonRpcSocketServer() {

	rpc.Register(new(Counter))

	l, err := net.Listen("tcp", ":3333")
	if err != nil {
		fmt.Printf("Listener tcp err: %s", err)
		return
	}

	for {
		fmt.Println("wating...")
		conn, err := l.Accept()
		if err != nil {
			fmt.Sprintf("accept connection err: %s\n", conn)
		}
		go jsonrpc.ServeConn(conn)
	}

}

func NewJsonRpcSocketClient() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:3333", 1000*1000*1000*30)
	if err != nil {
		fmt.Printf("create client err:%s\n", err)
		return
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)
	var reply int
	err = client.Call("Counter.Add", 10, &reply)

	fmt.Printf("reply: %s, err: %s\n", reply, err)

}
