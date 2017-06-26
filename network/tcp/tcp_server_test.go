package tcp

import "testing"

func TestTcpServer(t *testing.T) {
	DefaultServer()
}

// echo server
func TestTcpServerEcho(t *testing.T) {
	EchoServer()
}

func TestTcpClient(t *testing.T) {
	DefaultClient()
}
