package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panicf("cannot listen: %s", err)
	}
	log.Println("go404 listening on 8081")

	http.Serve(l, nil)

}
