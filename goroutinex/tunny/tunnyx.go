package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	pool := tunny.NewFunc(2, func(payload interface{}) interface{} {

		printHello()

		return ""
	})

	for i := 0; i < 10; i++ { // 控制并发量的例子
		wg.Add(1) // 注意：需要写到go fun外面
		go func() {
			defer wg.Done()
			pool.Process(printHello)
		}()
		fmt.Println("***")
	}
	wg.Wait()
	fmt.Println("all done")
}
func printHello() { // 无参数，无返回值
	fmt.Println("Hello!")
	time.Sleep(1 * time.Second)
}
