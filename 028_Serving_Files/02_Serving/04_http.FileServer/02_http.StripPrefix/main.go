package main

import (
	"io"
	"net/http"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", dog)
	//handle takes a route and the handler,
	// the StripPrefix strips the "/assets" off
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":7000", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/assets/card.jpg">`)
}
