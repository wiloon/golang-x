package main

import (
	"time"
	"fmt"

)

func main() {
	str := time.Now().Format(time.RFC1123Z)
	fmt.Println(str)

	str = time.Now().UTC().Format(time.RFC1123Z)
	fmt.Println(str)

	str = time.Now().Format(time.RFC3339)
	fmt.Println(str)

	format:="2006-01-02 15:04:05"
	fmt.Println(time.Now().Format(format))
}
