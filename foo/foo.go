package main

import (
	"fmt"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)
	http.ListenAndServe(":8000", n)
}
