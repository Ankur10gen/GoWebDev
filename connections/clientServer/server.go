package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
	fmt.Fprintln(conn, "You have successfully connected to the server")
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Connection Timed Out")
	}
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	defer conn.Close()

}
