package goroutinex

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Println(Goid())
	}()
	go func() {
		log.Println(Goid())
	}()
	log.Println(Goid())

	wg.Wait()
}
