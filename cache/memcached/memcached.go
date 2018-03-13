package main

import (
	"github.com/bradfitz/gomemcache/memcache"
	"fmt"
)

func main() {
	key := "foo"
	value := "value1"

	mc := memcache.New("192.168.36.47:11311", "192.168.36.48:11311")

	err := mc.Set(&memcache.Item{Key: key, Value: []byte(value)})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("set key:%v, value:%v\n", key, value)
	it, err := mc.Get(key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("value:", string(it.Value))

	fmt.Println("end.")
}
