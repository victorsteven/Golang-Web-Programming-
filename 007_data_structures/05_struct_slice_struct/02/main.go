package main

import (
	"log"
	"os"
	"text/template"
)

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	j := sage{
		Name:  "Jesus",
		Motto: "He is Love",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "Model T",
		Doors:        4,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corrolla",
		Doors:        4,
	}

	sages := []sage{j, g}
	cars := []car{f, c}

	//using an anonymous type where the underlying type is "struct"
	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
