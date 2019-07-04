package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	//read request
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			//headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	//request line
	m := strings.Fields(ln)[0] //method
	u := strings.Fields(ln)[1] //uri
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", u)

	//multiplexer
	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/about" {
		about(conn)
	}
	if m == "GET" && u == "/contact" {
		contact(conn)
	}
	if m == "GET" && u == "/apply" {
		apply(conn)
	}
	if m == "POST" && u == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head></head><meta charset="UTF-8"></title></head><body>
		<strong>Index</strong><br>
		<a href = "/">Index</a><br>
		<a href = "/about">About</a><br>
		<a href = "/contact">Contact</a><br>
		<a href = "/apply">Apply</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Conten-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
func about(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head></head><meta charset="UTF-8"></title></head><body>
		<strong>About</strong><br>
		<a href = "/">Index</a><br>
		<a href = "/about">About</a><br>
		<a href = "/contact">Contact</a><br>
		<a href = "/apply">Apply</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Conten-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head></head><meta charset="UTF-8"></title></head><body>
		<strong>Contact</strong><br>
		<a href = "/">Index</a><br>
		<a href = "/about">About</a><br>
		<a href = "/contact">Contact</a><br>
		<a href = "/apply">Apply</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Conten-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head></head><meta charset="UTF-8"></title></head><body>
		<strong>Apply</strong><br>
		<a href = "/">Index</a><br>
		<a href = "/about">About</a><br>
		<a href = "/contact">Contact</a><br>
		<a href = "/apply">Apply</a><br>
		<form method="post" action="/apply">
		<input type ="submit" value = "apply">
		</form>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Conten-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head></head><meta charset="UTF-8"></title></head><body>
		<strong>Apply Process</strong><br>
		<a href = "/">Index</a><br>
		<a href = "/about">About</a><br>
		<a href = "/contact">Contact</a><br>
		<a href = "/apply">Apply</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Conten-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

// //this is the responds to the server
// func respond(conn net.Conn) {
// 	body := `<!DOCTYPE html><html lang="en"><head></head><meta charset="UTF-8"></title></head><body>Hello Sir</body></html>`
// 	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
// 	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
// 	fmt.Fprint(conn, "Conten-Type: text/html\r\n")
// 	fmt.Fprint(conn, "\r\n")
// 	fmt.Fprintf(conn, body)
// }
