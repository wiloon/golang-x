package main

import (
	"net/http"
	"net/http/cgi"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler := new(cgi.Handler);
		handler.Path = "/usr/local/go/bin/go";
		script := "/Users/liujinlong/cgi-script" + r.URL.Path;
		fmt.Println(handler.Path);
		handler.Dir = "/Users/liujinlong/cgi-script";
		args := []string{"run", script};
		handler.Args = append(handler.Args, args...);
		fmt.Println(handler.Args);
		handler.ServeHTTP(w, r);
	});
	http.ListenAndServe(":8989", nil);
	select {} //阻塞进程
}
