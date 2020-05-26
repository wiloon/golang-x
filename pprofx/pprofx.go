package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for {
			log.Println(Add("https://github.com/foo"))
		}
	}()

	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}
