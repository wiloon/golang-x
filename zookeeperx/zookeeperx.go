package zookeeperx

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"fmt"
	"log"
)

const ROOT_PATH = "/"

type ZkNode struct {
	path  string
	value string
}

func (node ZkNode) getChildren(conn *zk.Conn) []ZkNode {
	children, _, err := conn.Children(node.path)
	if err != nil {
		panic(err)
	}

	nodes:=[len(children)]ZkNode
	for i, v := range children {
		log.Printf("%v, %v\n", i, v)
		nodes[i].path=v

	}
}

func foo() {
	connection, _, _ := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	defer connection.Close()

	root := ZkNode{path: ROOT_PATH}
	root.getChildren()

	children, _, err := connection.Children(ROOT_PATH)
	if err != nil {
		panic(err)
	}

}

func GetWithWatch() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	defer c.Close()

	if err != nil {
		panic(err)
	}

	children, stat, ch, err := c.ChildrenW("/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", children, stat)
	e := <-ch
	fmt.Printf("%+v\n", e)

}

func GetChildren() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	defer c.Close()

	if err != nil {
		panic(err)
	}

	children, _, err := c.Children(root)
	if err != nil {
		panic(err)
	}

	for i, v := range children {
		log.Printf("%v, %v\n", i, v)

	}

	fmt.Printf("%+v\n", children)

	fmt.Println(children[2])

	value, _, _ := c.Get("/" + children[2])
	fmt.Println(string(value))
}
