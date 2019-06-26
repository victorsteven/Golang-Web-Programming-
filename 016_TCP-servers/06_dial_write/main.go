package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//Fprintln takes a writer, and writes the string "i dialed you"
	fmt.Fprintln(conn, "i dialed you.")
}
