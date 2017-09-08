package tcp

import "testing"

//server
func TestTcpServer(t *testing.T) {
	DefaultServer()
}

//client
func TestTcpClient(t *testing.T) {
	DefaultClient()
}

// echo server
func TestTcpServerEcho(t *testing.T) {
	EchoServer()
}


