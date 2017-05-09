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

	acl := zk.WorldACL(zk.PermAll)
	conn.Create("/clipboard", []byte("data"), 0, acl)
	fmt.Println("create.")
	arr, state, e := conn.Get("/clipboard")
	if e != nil {

	}
	fmt.Println("value:", string(arr))
	fmt.Println("state:", state)

	conn.Set("/clipboard", []byte("data0"), 0)
	arrb, stateb, e := conn.Get("/clipboard")
	if e != nil {

	}
	fmt.Println("value:", string(arrb))
	fmt.Println("stateb:", stateb)
	//children, stat, ch, err := conn.ChildrenW("/")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("aaaa%+v %+v\n", children, stat)
	//e := <-ch
	//fmt.Printf("bbbb%+v\n", e)
}