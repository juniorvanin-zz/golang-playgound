package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handlerDefault)
	http.HandleFunc("/count", handlerCount)
	http.ListenAndServe("localhost:7000", nil)
}

func handlerCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Number of counts = %d", count)
	mu.Unlock()
}

func handlerDefault(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}
