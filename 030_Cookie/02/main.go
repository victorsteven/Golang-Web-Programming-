package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
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
	c1, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(res, "Your Cookie: #1", c1)

	c2, err := req.Cookie("general")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(res, "Your Cookie: #2", c2)

	c3, err := req.Cookie("specific")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(res, "Your Cookie: #3", c3)
}

func abundance(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "general",
		Value: "Some general",
	})
	http.SetCookie(res, &http.Cookie{
		Name:  "specific",
		Value: "Some specific",
	})

	fmt.Fprintln(res, "The other two cookies written")
}
