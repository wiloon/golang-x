package main

import (
	"time"
	"container/list"
)

func main() {

}

var bidList *list.List

func init() {
	bidList = list.New()
}

func Calculate(price float32, timestamp time.Time) {
	bidList.PushBack(price)
}
