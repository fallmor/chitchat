package main

import (
	"fmt"
	"net/http"
)

func handler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Bonjour Mor t es la %s", r.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
