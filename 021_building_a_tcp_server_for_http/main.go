package main

import (
	"fmt"
	"net/http"
)

//the underlying type is an int
type hotdog int

//Attaching a method to "hotdog" type
// the ServeHTTP method has the Handler interface
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	// it means any value of type "hotdog" is implicitly implementing the handler interface
	var d hotdog
	http.ListenAndServe(":7000", d)
}
