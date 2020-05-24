package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const address = ":8080"

func main() {
	log.Println("http mock server start, listen ", address)
	http.HandleFunc("/ping", SayHello)
	_ = http.ListenAndServe(address, nil)
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	response := `{"message":"pong"}`
	_, _ = w.Write([]byte(response))
	result, _ := ioutil.ReadAll(req.Body)
	_ = req.Body.Close()
	fmt.Printf("%s\n", result)
}
