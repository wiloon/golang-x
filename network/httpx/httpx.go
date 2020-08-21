package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGTERM)

	http.HandleFunc("/", hello)
	address := ":8080"
	log.Println("listening " + address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Println("err...")
		log.Println(err)
	}

	for s := range signals {
		if s == os.Interrupt || s == os.Kill || s == syscall.SIGTERM {
			break
		}
	}
	signal.Stop(signals)
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "hello w")
}
