package main

import (
	"fmt"
	"net/http"
)

// net/http servers use need functions serving as handlers that take a responsewriter and
// request as arguments to be wrapped by http.HandlerFunc adapter.

// prints hello
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

// prints all request headers into the response body
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		// headers is a slice for the name of the header so we print each header's slice element
		for _, value := range headers {
			fmt.Fprintf(w, "%v:%v\n", name, value)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	_ = http.ListenAndServe(":8080", nil)
}
