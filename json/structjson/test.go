package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Foo struct {
	Field0 string
	Field1 string
}

type Bar struct {
	Foo
	Field2 string
}

func main() {
	bar := Bar{
		Field2: "f2",
	}
	bar.Field0 = "f0"
	jb, _ := json.Marshal(bar)
	j := string(jb)
	fmt.Println(j)

	fmt.Println("---")
	foo := os.Stdout
	encoder := json.NewEncoder(foo)
	err := encoder.Encode(&bar)
	fmt.Println(err)
}
