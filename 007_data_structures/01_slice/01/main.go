package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	//this is the a composite literal of type string
	sages := []string{"Joy", "Peace", "Happiness"}

	err := tpl.Execute(os.Stdout, sages)

	if err != nil {
		log.Fatalln(err)
	}
}
