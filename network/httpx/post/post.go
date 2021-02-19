package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Server struct {
	ServerName string
	ServerIp   string
}

type ServerSlice struct {
	Server    []Server
	ServersID string
}

func main() {
	//post 第三个参数是io.reader interface
	//strings.NewReader  byte.NewReader bytes.NewBuffer  实现了read 方法

	resp, _ := http.Post(
		"http://baidu.com",
		"application/x-www-form-urlencoded",
		strings.NewReader("test"))

	defer resp.Body.Close()
	//io.Reader

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
