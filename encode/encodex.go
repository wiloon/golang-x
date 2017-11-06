package main

import (
	"fmt"
	"strconv"
	// "unicode/utf8"
)

func main() {
	rs := []rune("golang中文unicode编码")
	json := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			json += string(r)
		} else {
			json += "\\u"+strconv.FormatInt(int64(rint), 16) // json
		}
	}
	fmt.Printf("JSON: %s\n", json)

	r, _ := strconv.Unquote(`"` + json + `"`)

	fmt.Printf(r)


	//fmt.Printf(string(json))
}