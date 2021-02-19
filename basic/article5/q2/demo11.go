package main

import (
	"fmt"
)

var container = []string{"zero", "one", "two"}

func main() {
	fmt.Printf("The element is %q.\n", container[0])
	_, ok := interface{}(container).([]string)
	fmt.Println(ok)
	container := map[int]string{0: "zero_0", 1: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[0])
	_, ok = interface{}(container).([]string)

	fmt.Println(ok)
}
