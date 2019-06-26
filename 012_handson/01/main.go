package main

import (
	"html/template"
	"log"
	"os"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	AcaYear              string
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	years := []year{
		year{
			AcaYear: "2020-2021",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CHM-240", "Organic Chemistry", "50"},
					course{"PHY-250", "Principles of Physics", "40"},
					course{"COM-240", "Computer Science", "60"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CHM-240", "Organic Chemistry", "50"},
					course{"PHY-250", "Principles of Physics", "40"},
					course{"COM-240", "Computer Science", "60"},
				},
			},
		},

		year{
			AcaYear: "2022-2023",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CHM-240", "Organic Chemistry", "50"},
					course{"PHY-250", "Principles of Physics", "40"},
					course{"COM-240", "Computer Science", "60"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CHM-240", "Organic Chemistry", "50"},
					course{"PHY-250", "Principles of Physics", "40"},
					course{"COM-240", "Computer Science", "60"},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, years)
	if err != nil {
		log.Fatalln(err)
	}
}
