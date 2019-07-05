package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":7000", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, `<h1><a href = "/set">Set a cookie</a></h1>`)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "session",
		Value: "New value",
	})
	fmt.Fprintln(res, `<h1><a href="/read">read</a></h1>`)
}

func read(res http.ResponseWriter, req *http.Request) {

	c1, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return
	}
	// fmt.Fprintln(res, c1)
	fmt.Fprintf(res, `<h1>Your Cookie: <br>%v</h1><a href="/expire">expire</a><h1>`, c1)
}

func expire(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1 //delete cookie
	http.SetCookie(res, c)
	http.Redirect(res, req, "/", http.StatusSeeOther)

	//StatusSeeOther means i want to see the other resource, go get it more me
}
