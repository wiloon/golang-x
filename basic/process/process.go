package main

import (
	"fmt"
	"os"
)

func main() {
	name := "/usr/bin/ls"
	args := []string{"/"}
	attr := &os.ProcAttr{}
	proc, err := os.StartProcess(name, args, attr)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = proc.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
