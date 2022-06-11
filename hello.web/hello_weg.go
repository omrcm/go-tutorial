package main

import (
	"fmt"
	"net/http"
)

/*
	All Go programs start's running with main function. This function is called from main package
*/
func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})
	http.ListenAndServe(":8080", nil)
}
