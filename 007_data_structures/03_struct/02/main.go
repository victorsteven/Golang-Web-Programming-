package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type car struct {
	Name        string
	Description string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	honda := car{
		Name:        "Honda",
		Description: "This is a Japanese vehicle",
	}

	err := tpl.Execute(os.Stdout, honda)

	if err != nil {
		log.Fatalln(err)
	}
}
