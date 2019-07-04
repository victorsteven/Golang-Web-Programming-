package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

//below, the key is a string while the value is a function
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Name        string
	Description string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	//give me the first 3 ie: [0 1 2]
	s = s[:3]
	return s
}

func main() {
	s := sage{
		Name:  "Steven",
		Motto: "Wise guy",
	}

	b := sage{
		Name:  "Bola",
		Motto: "Good guy",
	}

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

	cars := []car{honda, mercedeze, ford}

	sages := []sage{s, b}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)

	if err != nil {
		log.Fatalln(err)
	}
}
