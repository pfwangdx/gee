package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// $ curl http://localhost:9999/
	http.HandleFunc("/", indexHandler)

	// curl http://localhost:9999/hello
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "url.path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintln(w, "header[%q] = %q\n", k, v)
	}
}
