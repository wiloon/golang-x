package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	log.Println("http mock server start, listen 80")
	http.HandleFunc("/foo", SayHello)
	http.ListenAndServe(":80", nil)
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	w.Write([]byte("{\"Status\":\"SUCCESS\"}"))
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	fmt.Printf("%s\n", result)
}
