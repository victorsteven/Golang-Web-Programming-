package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/card.jpg", dogPic)
	http.ListenAndServe(":7000", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html, charset=utf-8")

	io.WriteString(res, `
	<img src="/card.jpg">
	`)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	//below will give us a pointer to a file and an error
	f, err := os.Open("card.jpg")
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}

	// put the response into the file(f)
	// io.Copy(res, f)

	//or
	//We can serve a single file with ServeContent
	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)

}
