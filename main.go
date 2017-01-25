package main

import (
	"fmt"
	"net/http"
)

// handler(controller)を保持しておく
func main() {
	fmt.Println("The Server runs with http://localhost:8080/")
	http.HandleFunc("/", ReadHTML)
	http.HandleFunc("/control", sendResponse)
	http.ListenAndServe(":8080", nil)
}

