package zookeeperx

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"time"
	"fmt"
	"os"
)

const ROOT_PATH = "/k0"

type ZkNode struct {
	path  string
	value string
}

func (node ZkNode) toString() string {
	return node.path + "=" + node.value
}
func (node ZkNode) getChildren(conn *zk.Conn) []ZkNode {
	parentPath := node.path
	log.Println("get children, path", node)
	children, _, err := conn.Children(parentPath)
	if err != nil {
		panic(err)
	}
	log.Println("children:", children)

	nodes := []ZkNode{}
	subChildren := []ZkNode{}
	if len(children) == 0 {
		node.value = node.getValue(conn)
		nodes = append(nodes, node)
		log.Println("append node:", node)
	} else {

		for i, v := range children {
			log.Printf("child %v, %v\n", i, v)
			childNode := ZkNode{path: parentPath + "/" + v}
			log.Println("child node created:", childNode)
			subChildren = childNode.getChildren(conn)
			nodes = append(nodes, subChildren...)
		}

	}
	return nodes
}

func (node ZkNode) getValue(conn *zk.Conn) string {
	b, _, _ := conn.Get(node.path)
	value := string(b)
	log.Println("get value:", value)
	return value
}

func foo() {
	connection, _, _ := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	defer connection.Close()

	root := ZkNode{path: ROOT_PATH}
	children := root.getChildren(connection)

	file, err := os.Create("export.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	for i, v := range children {
		log.Printf("%v,%v\n", i, v)
		_, err := file.WriteString(v.toString()+"\n")
		if err != nil {
			panic(err)
		}
	}
	file.Sync()
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

	children, _, err := c.Children(ROOT_PATH)
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
