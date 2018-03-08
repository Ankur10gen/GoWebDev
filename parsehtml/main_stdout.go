package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("templatefile.gohtml")
	if err != nil {
		log.Fatalln("Template file not found",err)
	}
	//reusing err variable to hold error from Execute function
	//Execute here is applying the parsed template in tpl to nil. Writing to stdout in this case
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
