package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var name string

	if len(os.Args) > 1 {
		name = os.Args[1]
	} else {
		name = "Ankur"
	}

	str := fmt.Sprint(
		`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>
		stringtohtml
		</title>
		<body> My name is `,
		name,
		` 
		</body>
		</html>
		`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Unable to create file")
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))
}
