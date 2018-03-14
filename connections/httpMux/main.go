package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)

	//respond(conn)
}

func request(conn net.Conn) {
	//GET /hello/world HTTP/1.1
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m, u := strings.Fields(ln)[0], strings.Fields(ln)[1]
			fmt.Println("***METHOD***", m)
			if m == "GET" && u == "/" {
				home(conn)
			}

			if m == "GET" && u == "/aboutus" {
				aboutus(conn)
			}

			if m == "GET" && u == "/contactus" {
				contactus(conn)
			}

			if m == "POST" && u == "/contactus" {
				home(conn)
			}
		}
		if ln == "" {
			//headers are done
			break
		}
		i++
	}
}
func home(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charet="UTF-8">
	<title></title>
	</head>
	<body>
	<strong>This is HOME</strong>
	</body>
	</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") //most important
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func aboutus(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charet="UTF-8">
	<title></title>
	</head>
	<body>
	<strong>About Us</strong>
	</body>
	</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") //most important
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contactus(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charet="UTF-8">
	<title></title>
	</head>
	<body>
	<strong>Contact Us</strong>
	<br>
	<form method="POST" action="/contactus">
	<input type="submit" value="apply">
	</form>
	</body>
	</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") //most important
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
