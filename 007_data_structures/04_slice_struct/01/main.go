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
	mercedeze := car{
		Name:        "Mercedeze",
		Description: "This is a German vehicle",
	}
	ford := car{
		Name:        "Ford",
		Description: "This is an American vehicle",
	}

	//defining the slice of struct
	cars := []car{honda, mercedeze, ford}

	err := tpl.Execute(os.Stdout, cars)

	if err != nil {
		log.Fatalln(err)
	}
}
