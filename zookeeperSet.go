package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	conn, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	arr, state, e := conn.Get("/clipboard")
	if e != nil {

	}
	fmt.Println("value:", string(arr))
	fmt.Println("state:", state)
	fmt.Println("state.version:", state.Version)
	conn.Set("/clipboard", []byte("connected4"), state.Version)
	fmt.Println("done.")
}