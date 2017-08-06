package zookeeperx

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"time"
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
)

type ZkNode struct {
	path    string
	value   string
	version int32
}

func Delete(path string) {
	connection, _, _ := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	defer connection.Close()

	node := ZkNode{path: path}
	node.Delete(connection)
}

func (node ZkNode) Delete(conn *zk.Conn) {
	log.Println("deleting node:", node)
	children := node.getChildren(conn)

	if len(children) == 0 {
		err := conn.Delete(node.path, node.version)
		if err != nil {
			panic(err)
		}
		log.Printf("node deleted, path:%v,value:%v", node.path, node.value)
	} else {
		for _, v := range children {
			v.Delete(conn)

		}

	}

}

func (node ZkNode) createNode(conn *zk.Conn) {
	conn.Create(node.path, []byte(node.value), 0, zk.WorldACL(zk.PermAll))
}

func (node ZkNode) toString() string {
	return node.path + "=" + node.value
}

func (node ZkNode) hasChildren(conn *zk.Conn) bool {
	children, _, err := conn.Children(node.path)
	if err != nil {
		panic(err)
	}
	return !(len(children) == 0)
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

	if len(children) > 0 {
		for i, v := range children {
			log.Printf("child %v, %v\n", i, v)
			childNode := ZkNode{path: parentPath + "/" + v}
			nodes = append(nodes, childNode)
			log.Println("child found:", childNode)

			subChildren := childNode.getSubChildren(conn)
			nodes = append(nodes, subChildren...)
			log.Println("merge sub child:", subChildren)
		}

	}
	log.Printf("%v children:%v", parentPath, nodes)
	return nodes
}

func (node ZkNode) getSubChildren(conn *zk.Conn) []ZkNode {
	parentPath := node.path
	log.Println("get sub children, path", node)
	children, _, err := conn.Children(parentPath)
	if err != nil {
		panic(err)
	}
	log.Println("children:", children)

	nodes := []ZkNode{}
	subChildren := []ZkNode{}

	if len(children) > 0 {
		for i, v := range children {
			log.Printf("child %v, %v\n", i, v)
			childNode := ZkNode{path: parentPath + "/" + v}
			log.Println("child found:", childNode)
			subChildren = childNode.getSubChildren(conn)
			nodes = append(nodes, subChildren...)
		}

	} else {
		node.value, node.version = node.getValue(conn)
		nodes = append(nodes, node)
		log.Println("collect child node:", node)
	}

	return nodes
}

func (node ZkNode) getValue(conn *zk.Conn) (string, int32) {
	b, stat, _ := conn.Get(node.path)
	value := string(b)

	log.Printf("get value:%v,stat:%v", value, stat)
	return value, stat.Version
}

func Export(path string) {
	connection, _, _ := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	defer connection.Close()

	root := ZkNode{path: path}
	children := root.getChildren(connection)

	file, err := os.Create("export.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	for i, v := range children {
		log.Printf("children list: %v,%v\n", i, v)
		_, err := file.WriteString(v.toString() + "\n")
		if err != nil {
			panic(err)
		}
	}
	file.Sync()
}

func importFromFile() {

	connection, _, _ := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	defer connection.Close()

	filePath := "export.txt"
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	for {
		n, _, err := r.ReadLine()
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == len(n) {
			break
		}
		line := string(n)
		arr := strings.Split(line, "=")
		node := ZkNode{path: arr[0], value: arr[1]}
		node.createNode(connection)

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

func GetChildren(path string) {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	defer c.Close()

	if err != nil {
		panic(err)
	}

	children, _, err := c.Children(path)
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
