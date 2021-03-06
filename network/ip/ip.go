package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"log"
)

var get_ip = flag.String("get_ip", "", "external|internal")

func main() {
	fmt.Println("Usage of ./getmyip --get_ip=(external|internal)")
	flag.Parse()
	//if *get_ip == "external" {
	get_external()
	//}

	//if *get_ip == "internal" {
	get_internal()
	//}

}

func get_external() {
	url:="http://members.3322.org/dyndns/getip"
	//url="http://ident.me/"
	resp, err := http.Get(url)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	log.Println("ip:",resp.Body)

	//os.Exit(0)
}

func get_internal() {
	addrs, err := net.InterfaceAddrs()

	log.Println(len(addrs))
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		log.Println(a)
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
	os.Exit(0)
}
