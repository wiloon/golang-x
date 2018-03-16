package main

import (
	"github.com/robfig/cron"
	"log"
)

func main() {
	i := 0
	c := cron.New()
	spec := "0 */1 * * * *"
	c.AddFunc(spec, func() {
		i++
		log.Println("start", i)
	})
	c.Start()
	select{} //阻塞主线程不退出

}
