package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
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
	// io.WriteString(res, "hello Steven")
	//parse the file and give us back a pointer to a template and an error
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	//Execute the template
	err = tpl.ExecuteTemplate(res, "index.gohtml", "Hello Steven")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello bar")
}
