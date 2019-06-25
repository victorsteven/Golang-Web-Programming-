package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))

}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)

	if err != nil {
		log.Fatalln(err)
	}
}
