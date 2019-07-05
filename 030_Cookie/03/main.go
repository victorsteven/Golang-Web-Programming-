package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":7000", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}
	//strconv: string convert
	//Atoi: ASCII to Int
	count, _ := strconv.Atoi(cookie.Value)
	count++
	//Itoa: Int to ASCII
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)

	io.WriteString(res, cookie.Value)
}
