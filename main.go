package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func status(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "healthy\n")
}

func main() {
	fmt.Println("Starting Test Release Please Server")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/status", status)
	http.ListenAndServe(":8090", nil)
}
