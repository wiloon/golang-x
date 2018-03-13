package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Println("Hello, 世界")
	path := "/tmp/to/create3"

	// check
	if _, err := os.Stat(path); err == nil {
		fmt.Println("path exists 1", path)
	} else {
		fmt.Println("path not exists ", path)
		err := os.MkdirAll(path, 0711)

		if err != nil {
			log.Println("Error creating directory")
			log.Println(err)
			return
		}
	}

	// check again
	if _, err := os.Stat(path); err == nil {
		fmt.Println("path exists 2", path)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}