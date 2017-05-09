package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	response,_ := http.Get("http://www.baidu.com")
	defer response.Body.Close()
	body,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
