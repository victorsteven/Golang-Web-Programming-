package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

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

//from package "time" call type "Time"
func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())

	if err != nil {
		log.Fatalln(err)
	}
}
