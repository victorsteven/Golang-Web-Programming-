package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServerHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "doggy")
	case "/cat":
		io.WriteString(res, "kitty")
	}
}

func main() {
	//d is of type "hotdog", type of handler and an underying type of int
	var d hotdog
	http.ListenAndServe(":7000", d)
}
