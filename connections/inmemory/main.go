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
		defer conn.Close()
	}
}

func handle(conn net.Conn) {
	fmt.Fprintln(conn, "You are using a silly in memory database")
	fmt.Fprintln(conn, `Use commands:
		SET key value - to set the key to a value
		GET key - to get the value
		DEL key - to delete the key & value`)

	inMemBuffer := make(map[string]string)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")

		switch cmd[0] {
		case "SET":
			inMemBuffer[cmd[1]] = cmd[2]
			fmt.Fprintf(conn, "Key %s set to value %s \n", cmd[1], cmd[2])
		case "GET":
			val, exists := inMemBuffer[cmd[1]]
			if !exists {
				fmt.Fprintln(conn, "Value doesn't exist.")
			} else {
				fmt.Fprintf(conn, "Key %s has value %s \n", cmd[1], val)
			}
		case "DEL":
			delete(inMemBuffer, cmd[1])
			fmt.Fprintf(conn, "Key %s Deleted \n", cmd[1])

		default:
			fmt.Fprintln(conn, "Wrong command.")
			fmt.Fprintln(conn, `Use commands:
			SET key value - to set the key to a value
			GET key - to get the value
			DEL key - to delete the key & value`)
		}
	}
}
