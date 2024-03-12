package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

var originServerURL *url.URL
var err error

func main() {

	// define origin server URL
	originServerURL, err = url.Parse("http://127.0.0.1:8081")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	http.HandleFunc("/", reverseProxyHandler)
	http.HandleFunc("/ping", pingHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func reverseProxyHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("[reverse proxy server] received request at: %s\n", time.Now())

	// set req Host, URL and Request URI to forward a request to the origin server
	req.Host = originServerURL.Host
	req.URL.Host = originServerURL.Host
	req.URL.Scheme = originServerURL.Scheme
	req.RequestURI = ""

	// save the response from the origin server
	originServerResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(rw, err)
		return
	}

	// return response to the client
	rw.WriteHeader(http.StatusOK)
	io.Copy(rw, originServerResponse.Body)
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("[origin server] received ping request at: %s\n", time.Now())
	_, _ = fmt.Fprint(rw, "pong\n")
}
