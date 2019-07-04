package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		log.Fatal(err)
	}
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close()

	// err = tpl.Execute(os.Stdout, nil)
	//Execute takes a Writer and data
	err = tpl.Execute(nf, nil)

	if err != nil {
		log.Fatalln(err)
	}
}
