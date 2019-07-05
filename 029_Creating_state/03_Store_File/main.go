package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

// var tpl *template.Template

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":7000", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	var s string

	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		// http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nfile:", f, "\nheader:", h)

		//read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		//store on server
		dst, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(res, "index.gohtml", s)

	// res.Header().Set("Content-Type", "text/html; charset=utf-8")
	// io.WriteString(res, `
	// 	<form method="POST" enctype="multipart/form-data">
	// 	<input type="file" name="q">
	// 	<input type="submit">
	// 	</form>
	// 	<br>
	// `+s)
}
