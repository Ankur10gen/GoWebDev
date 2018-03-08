package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	//Save error checking for template parsing as Must checks it for template pointers
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "a.gohtml", nil)
	if err != nil {
		log.Fatalln("Failed while executing template ", err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "b.gohtml", nil)
	if err != nil {
		log.Fatalln("Failed while executing template ", err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "c.gohtml", nil)
	if err != nil {
		log.Fatalln("Failed while executing template ", err)
	}

}
