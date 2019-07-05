package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":7000", nil)
}
func set(res http.ResponseWriter, req *http.Request) {
	//for the composite literal below, the type is "&http.Cookie"
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "Some value",
	})
	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK YOUR BROWSER")
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(res, "Your Cookie:", c)
}
