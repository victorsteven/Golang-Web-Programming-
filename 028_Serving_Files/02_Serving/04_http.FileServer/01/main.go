package main

import (
	"io"
	"net/http"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir(".")))
	// http.HandleFunc("/dog", dog)
	// http.ListenAndServe(":7000", nil)

	//or, we here we will be serving main.go, and we dont want this
	http.ListenAndServe(":7000", http.FileServer(http.Dir(".")))

}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="card.jpg">`)
}
