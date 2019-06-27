package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Steven-Key", "this is from Steven")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want ti the func</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":7000", d)
}
