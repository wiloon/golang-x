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

func Calculate(bid float64, timestamp time.Time) {

	bidList.PushBack(price{bid, timestamp})
}

type price struct {
	bid       float64
	timestamp time.Time
}
