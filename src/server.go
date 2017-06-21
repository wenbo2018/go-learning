package main

import (
	"net/http"
	"log"
	"sync"
)

var mu sync.Mutex
var count int
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


