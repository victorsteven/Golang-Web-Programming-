package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to computer programming", "4"},
				course{"CSCI-130", "Second Introduction to computer programming", "5"},
				course{"CSCI-140", "Third Introduction to computer programming", "10"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-40", "Advanced Introduction to computer programming", "10"},
				course{"CSCI-130", "Advanced Second Introduction to computer programming", "51"},
				course{"CSCI-140", "Advanced Third Introduction to computer programming", "60"},
			},
		},
	}
	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
