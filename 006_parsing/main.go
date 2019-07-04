package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	//using Must, we automatically check our errors, so no need of assigning an error variable
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "This is a string we parsed here")

	if err != nil {
		log.Fatalln(err)
	}
}
