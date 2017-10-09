package main

import (
	"net"
	"log"
)

func main() {
	services := []string{
		"localhost:2181",
		"localhost:7000",
		"localhost:7001",
		"localhost:7002",
		"localhost:7003",
		"localhost:7004",
		"localhost:7005",
		"localhost:8086",
		"localhost:3306",
	}

	for _, address := range services {
		socket_ip(address)
	}
}

func socket_ip(address string) bool {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", address) //转换IP格式

	conn, err := net.DialTCP("tcp", nil, tcpAddr) //查看是否连接成功
	if err != nil {
		log.Printf("%s, inactive\r\n", address)
		return false

	}
	defer conn.Close()
	log.Printf("%s, active\r\n", address)
	return true
}
