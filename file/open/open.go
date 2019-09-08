package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "/data/rssx/logs/rssx.log"
	//filePath = "/tmp/foo.log"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Failed to log to file, using default stderr,error:", err)
	}

	defer file.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := file.Write(d2)
	check(err)
	file.WriteString("foo")
	fmt.Printf("wrote %d bytes\n", n2)
	fmt.Println("end.")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
