package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":7000", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `
	<!--This will not serve ->
	<img src="/ball.png">
	`)
}
