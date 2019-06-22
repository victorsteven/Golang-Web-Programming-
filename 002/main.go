package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	// name := "Steven Victor"
	// tp := `
	// 	<!DOCTYPE html>
	// 	<html lang="en">
	// 	<head>
	// 		<meta charset="UTF-8">
	// 		<title>Hello Dear!</title>
	// 	</head>
	// 	<body>
	// 		<h1>` + name + `</h1>
	// 	</body>
	// 	</html>
	// `
	// fmt.Println(tp)

	// name := os.Args[1]
	name := "Agu Mark"
	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Hello Dear!</title>
		</head>
		<body>
			<h1>` + name + `</h1>
		</body>
		</html>
	`)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating the file", err)
	}
	defer nf.Close()

	//takes in a writer and a reader
	io.Copy(nf, strings.NewReader(str))
}
