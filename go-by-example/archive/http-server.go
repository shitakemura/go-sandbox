package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, rep *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
