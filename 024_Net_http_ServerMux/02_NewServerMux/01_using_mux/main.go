package main

import (
	"io"
	"net/http"
)

type hotdog int

//creating a handler of type hotdog
func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog")
}

type hotcat int

//creating a handler of type hotcat
func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux() //this returns the pointer *ServeMux
	//so the value "mux" is of type "pointer to a ServeMux" (*ServeMux")
	//the first parameter is the pattern, second is the handler
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":7000", mux)

}
