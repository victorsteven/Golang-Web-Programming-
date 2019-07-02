package main

import (
	"io"
	"net/http"
)

func main() {
	//HandleFunc() is used to add routes, the second argument in the function is a route

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/me/", myName)
	http.ListenAndServe(":7000", nil) //the defualt ServeMux is used
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func myName(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "hello Steven")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello bar")
}
