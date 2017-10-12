package main

import (
	"log"
	"fmt"
	"strings"
	"encoding/json"
)

type type0 struct {
	F0 string
	F1 string
}

func main() {
	log.Println(fmt.Sprintf("%013d", 1))

	t0 := type0{"a", "b\nc"}
	str, _ := json.Marshal(t0)
	strNew := string(str)

	log.Println("new:", strNew)
	newMsg := strings.Replace(strNew, "\n", "<br>", -1)
	log.Println("msg:", newMsg)

	c := "{\"F0\":\"a\",\"F1\":\"b\nc\"}"
	log.Println("c:", c)
	log.Println(strings.Replace(c, "\n", "<br>", -1))
}
