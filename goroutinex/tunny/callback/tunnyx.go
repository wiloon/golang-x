package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	pool2 := tunny.NewCallback(2)

	for i := 0; i < 10; i++ { // 控制并发量的例子
		wg.Add(1) // 注意：需要写到go fun外面
		go func() {
			defer wg.Done()
			pool2.Process(printHello)
		}()
	}
	wg.Wait()
	fmt.Println("all done")
}
func printHello() { // 无参数，无返回值
	fmt.Println("Hello!")
	time.Sleep(1 * time.Second)
}
