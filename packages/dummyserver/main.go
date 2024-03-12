package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ping", pingHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func rootHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("[dummy server] received root request at: %s\n", time.Now())
	_, _ = fmt.Fprint(rw, "Hello World!\n")
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("[dummy server] received ping request at: %s\n", time.Now())
	_, _ = fmt.Fprint(rw, "pong\n")
}
