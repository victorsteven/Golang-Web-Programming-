package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/card.jpg", dogPic)
	http.ListenAndServe(":7000", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html, charset=utf-8")

	io.WriteString(res, `
	<img src="/card.jpg">
	`)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "card.jpg")

}
