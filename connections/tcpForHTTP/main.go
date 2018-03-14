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

	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD***", m)
		}
		if ln == "" {
			//headers are done
			break
		}
		//fmt.Println(strings.Fields(ln)[0])
		//Used referer instead of GET URL
		if strings.Fields(ln)[0] == "Referer:" {
			fmt.Printf("The requested URL is %s\n", strings.Fields(ln)[1])
		}
		i++
	}
}
func respond(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charet="UTF-8">
	<title></title>
	</head>
	<body>
	<strong>Hello World</strong>
	</body>
	</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") //most important
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
