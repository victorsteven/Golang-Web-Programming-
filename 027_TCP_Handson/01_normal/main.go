package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":7000")

	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(conn, "I see the connection")

		conn.Close()
	}
}
