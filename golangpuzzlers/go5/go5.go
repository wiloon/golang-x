package main

import "fmt"

var block = "package"

func main() {
	block := "function"
	{
		foo()
		block := 3
		fmt.Printf("The block is %v.\n", block)
	}
	foo()
	fmt.Printf("The block is %s.\n", block)
}
func foo() {
	fmt.Printf("The block is %s.\n", block)
}
