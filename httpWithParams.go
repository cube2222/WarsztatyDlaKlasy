package main

import (
	"net/http"
	"net/url"
	"fmt"
	"strconv"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		vars, _ := url.ParseQuery(r.URL.RawQuery)
		x, _ := strconv.Atoi(vars["x"][0])
		y, _ := strconv.Atoi(vars["y"][0])
		fmt.Fprintln(w, "Hello, ", x + y)
	})
	http.ListenAndServe(":80", nil)


}
