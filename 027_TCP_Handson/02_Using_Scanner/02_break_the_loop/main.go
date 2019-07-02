package main

import (
	"bufio"
	"fmt"
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
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				fmt.Println("this is the end of the http request headers")
				break
			}
		}
		// defer conn.Close()

		fmt.Println("Code got here") //this is because, we are scanning an open stream that have no end
		io.WriteString(conn, "I see the connection")

		conn.Close()
	}
}
