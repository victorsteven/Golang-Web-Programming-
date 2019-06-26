package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	//Eternally loop through the listener

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Conn timeout")
	}
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I have it: %s\n", ln)
	}

	defer conn.Close()

	fmt.Println("We are not getting here")

}
