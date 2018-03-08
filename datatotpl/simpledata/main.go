package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	//Returns pointer to template after parsing and error checking
	tpl = template.Must(template.ParseGlob("./*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "templatefile.gohtml", "Ankur")
	if err != nil {
		log.Fatalln("Wasn't able to execute the template ", err)
	}
}
