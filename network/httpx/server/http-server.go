package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

const address = ":3000"

func main() {
	log.Println("http mock server start, listen ", address)
	http.HandleFunc("/foo", SayHello)
	http.ListenAndServe(address, nil)
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Write([]byte("{\"Status\":\"SUCCESS\"}"))
	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	fmt.Printf("%s\n", result)
}
