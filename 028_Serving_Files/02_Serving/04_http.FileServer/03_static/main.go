package main

import "net/http"

func main() {
	//this will serve everything in the directory specified
	http.ListenAndServe(":7000", http.FileServer(http.Dir(".")))
}
