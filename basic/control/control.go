package main

import (
	"time"
	"log"
)

func main() {
	for {
		log.Println("foo")
		time.Sleep(3 * time.Second)
	}
}

func test0() {
	log.Println("start")
	for false {
		time.Sleep(time.Second * 1)
		log.Println("***")
	}
	log.Println("end")
}

func test1() {

	log.Println("start")
	index := 0
	stop := false
	for !stop {
		time.Sleep(time.Second * 1)
		log.Println("index:", index)
		if index == 3 {
			stop = true
		}
		index++
	}
	log.Println("end")
}
