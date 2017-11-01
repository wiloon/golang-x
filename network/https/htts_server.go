package main

import (
	"net/http"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	//fmt.Fprintf(w, "This is an example server.\n")
	//io.WriteString(w, "This is an example server.\n")
	log.Println("hello https:", req)
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":8443", "/etc/nginx/ssl/ssl.crt", "/etc/nginx/ssl/ssl.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
