package main

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"log"
)

func main() {

	//跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)

	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}

	c.Get("https://baidu.com")

}

func https0() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", "220.181.57.216:443", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}
	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}
	println(string(buf[:n]))
}
