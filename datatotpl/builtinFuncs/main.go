package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	cars := []string{"i10", "i20", "dzire"}
	//cars := map[string]string{"i10": "hatchback", "i20": "hatchback", "duster": "suv"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", cars)
	if err != nil {
		log.Fatalln("Couldn't execute template", err)
	}
}
