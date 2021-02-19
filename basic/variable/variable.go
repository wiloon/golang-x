package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var s string
	fmt.Println(s)

	var s1 = "foo"
	fmt.Println(s1)

	s2 := "bar"
	fmt.Println(s2)

	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n")

	fmt.Println(n, err)

}
