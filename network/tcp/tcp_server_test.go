package tcp

import "testing"

func TestTcpServer(t *testing.T) {
	DefaultServer()
}

func TestTcpServerEcho(t *testing.T) {
	EchoServer()
}

func TestTcpClient(t *testing.T) {
	DefaultClient()
}

func TestTcpClientEcho(t *testing.T) {
	EchoClient()
}
