package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

func main() {
	defaultTestStr := `
{
    "field0": "value0",
    "field1": "value1"
}
`
	jsonStr := flag.String("json-string", defaultTestStr, "")
	flag.Parse()
	m := make(map[string]string)
	_ = json.Unmarshal([]byte(*jsonStr), &m)
	fmt.Println(m)
	fmt.Println("params:=make(map[string]string")
	for k, v := range m {
		fmt.Printf("params[\"%v\"]=\"%v\"\n", k, v)
	}
}
