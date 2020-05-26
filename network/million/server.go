package main

import (
	log "github.com/cihub/seelog"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {
	ln, err := net.Listen("tcp", ":8972")
	if err != nil {
		panic(err)
	}
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Errorf("pprof failed: %v", err)
		}
	}()
	var connections []net.Conn
	defer func() {
		for _, conn := range connections {
			conn.Close()
		}
	}()
	for {
		conn, e := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Infof("accept temp err: %v", ne)
				continue
			}
			log.Infof("accept err: %v", e)
			return
		}
		go handleConn(conn)
		connections = append(connections, conn)
		if len(connections)%100 == 0 {
			log.Infof("total number of connections: %v", len(connections))
		}
	}
}
func handleConn(conn net.Conn) {
	io.Copy(ioutil.Discard, conn)
}
