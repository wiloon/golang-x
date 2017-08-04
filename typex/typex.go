package typex

import "log"

type Node struct {
	f0 string
	v1 string
}

func foo() {
	n0 := Node{f0: "v0"}

	log.Println("1:", n0)
	log.Println("2:", &n0)

	n0.fun0()
	log.Println("3:", n0)
}

func (n Node) fun0() {
	n.f0 = "v00"
}

func bar() {

	n0 := new(Node)
	n0.f0 = "v0"
	n0.fun0()
	log.Println("3:", *n0)
}
