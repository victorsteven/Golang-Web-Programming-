package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":7000", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(res, "Awesome me: "+v)
}
