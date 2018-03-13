package main

import (
	"os"

	"log"
	"fmt"
)

func main() {
	file, err := os.OpenFile("/tmp/test.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Failed to log to file, using default stderr,error:", err)
	}

	defer file.Close()
	//你可以写入你想写入的字节切片
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := file.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
	fmt.Println("end.")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
