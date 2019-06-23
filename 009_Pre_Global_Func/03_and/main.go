package main

import (
	"html/template"
	"log"
	"os"
)

type user struct {
	Name  string
	Motto string
	Admin bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	u1 := user{
		Name:  "Steven",
		Motto: "Make sure you know what enters for you",
		Admin: false,
	}
	u2 := user{
		Name:  "Aloe",
		Motto: "Medicinal food",
		Admin: true,
	}
	u3 := user{
		Name:  "",
		Motto: "Anonymous",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
