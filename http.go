package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	})
	http.ListenAndServe(":80", nil)
}
