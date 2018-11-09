package main

import (
	"fmt"
	"log"
)

func main() {
	c := make(chan int)
	defer close(c)
	go func() {
		log.Println("write c")
		c <- 3 + 4
	}()
	log.Println("reading c")
	i := <-c
	fmt.Println(i)
}
