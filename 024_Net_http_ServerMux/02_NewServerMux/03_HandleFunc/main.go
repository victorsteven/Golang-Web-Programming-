package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat")
}

func main() {
	// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	//dont forget that ListenAndServe takes a handler interface
	http.ListenAndServe(":7000", nil)

}
