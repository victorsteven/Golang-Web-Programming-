// package main

// func handle(conn net.Conn) {
// 	scanner := bufio.NewScanner(conn)

// 	for scanner.Scan() {
// 		ln := strings.ToLower(scanner.Text())
// 		bs := []byte(ln)
// 		// r := rot13(bs)

// 		fmt.Fprintf(conn, "%s - %s\n\n", ln, r)
// 	}
// }

// func rot13(bs []byte) []byte {
// 	var r13 = make([]byte, len(bs))
// 	for i, v := range bs {
// 		if v <= 109 {
// 			r13[i] = v + 13
// 		} else {
// 			r13[i] = v - 13
// 		}
// 	}
// }
