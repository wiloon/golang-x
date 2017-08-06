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
	leaf    bool
}

func (node ZkNode) exist(conn *zk.Conn) bool {
	exist, _, _ := conn.Exists(node.path)
	return exist
}
func (node ZkNode) isRoot() bool {
	return strings.LastIndex(node.path, "/") == 0
}
func Delete(path string) {
	connection, _, _ := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	defer connection.Close()

	node := ZkNode{path: path}
	_, node.version = node.getValue(connection)

	children := node.getChildren(connection)
	GetValue(children)

	for _, v := range children {
		if !v.hasChildren(connection) {
			v.Delete(connection)
		}
	}
	node.Delete(connection)
}

func GetValue(nodes map[string]ZkNode) {
	log.Println("get value:", nodes)
	connection, _, _ := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	defer connection.Close()

	for k, v := range nodes {
		v.value, v.version = v.getValue(connection)
		nodes[k] = v
	}
}

func (node ZkNode) Delete(conn *zk.Conn) {
	log.Println("deleting node:", node)
	err := conn.Delete(node.path, node.version)
	if err != nil {
		panic(err)
	}
	log.Printf("node deleted, path:%v,value:%v", node.path, node.value)

}

func (node ZkNode) CreateNode(conn *zk.Conn) {
	log.Println("create node:", node)

	if node.exist(conn) {
		log.Println("node exist:", node)
	} else {
		index := strings.LastIndex(node.path, "/")
		if index > 0 {
			parentNode := ZkNode{path: node.path[0:index]}
			if parentNode.exist(conn) {
				conn.Create(node.path, []byte(node.value), 0, zk.WorldACL(zk.PermAll))
			} else {
				parentNode.CreateNode(conn)
			}

		} else {
			conn.Create(node.path, []byte(node.value), 0, zk.WorldACL(zk.PermAll))
		}
	}

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

func (node ZkNode) getChildren(conn *zk.Conn) map[string]ZkNode {
	parentPath := node.path
	log.Println("get children, path", node)
	children, _, err := conn.Children(parentPath)
	if err != nil {
		panic(err)
	}
	log.Println("children:", children)

	nodes := map[string]ZkNode{}

	if len(children) > 0 {
		for i, v := range children {
			log.Printf("child %v, %v\n", i, v)
			childNode := ZkNode{path: parentPath + "/" + v}
			nodes[childNode.path] = childNode
			log.Println("child found:", childNode)

			subChildren := childNode.getChildren(conn)
			for k, v := range subChildren {
				nodes[k] = v
			}
			log.Println("merge sub child:", subChildren)
		}

	} else if !node.isRoot() {
		node.leaf = true
		nodes[node.path] = node
		log.Println("collect child node:", node)
	}
	log.Printf("%v children:%v", parentPath, nodes)
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
	GetValue(children)

	file, err := os.Create("export.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	for k, v := range children {
		log.Printf("children list: key:%v,value:%v\n", k, v)
		_, err := file.WriteString(v.toString() + "\n")
		if err != nil {
			panic(err)
		}
	}

	file.Sync()
}

func ImportFromFile() {

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
		log.Println("line:", line)
		arr := strings.Split(line, "=")
		node := ZkNode{path: arr[0], value: arr[1]}
		node.CreateNode(connection)

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
