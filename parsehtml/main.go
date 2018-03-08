package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	//Parsing template file which returns a pointer to template

	tpl, err := template.ParseFiles("templatefile.gohtml")
	if err != nil {
		log.Fatalln("File was not found ", err)
	}

	//Create a file on os to hold the response from template execution

	outputfile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Unable to create file ", err)
	}

	//Use the template execute and pass the index.html as writer and apply on nil object
	//Execute returns only error (if any) and writes to the outputfile/a writer interface

	err = tpl.Execute(outputfile, nil)
	if err != nil {
		log.Fatalln("Couldn't execute the template")
	}

	tpl, err = template.ParseFiles("templatefile.gohtml_2", "templatefile.gohtml_1")
	if err != nil {
		log.Fatalln("couldn't find files ", err)
	}

	err = tpl.ExecuteTemplate(outputfile, "templatefile.gohtml_2", nil)
	if err != nil {
		log.Fatalln("Couldn't execute the template ", err)
	}

	tpl, err = tpl.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatalln("No errors found in directory ", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "a.gohtml", nil)
	if err != nil {
		log.Fatalln("Couldn't execute the template ", err)
	}
}
