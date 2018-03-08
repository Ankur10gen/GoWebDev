package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./*.gohtml*"))
}

func main() {

	listOfExecutives := []string{"Dev Ittycheria", "Eliot Horowitz", "Michael Gordon"}

	//Accessing the elements from list

	err := tpl.ExecuteTemplate(os.Stdout, "templatefile.gohtml", listOfExecutives)
	if err != nil {
		log.Fatalln("Oops! Wasn't able to execute template", err)
	}

	//Accessing indexes & elements

	err = tpl.ExecuteTemplate(os.Stdout, "templatefile.gohtml_1", listOfExecutives)
	if err != nil {
		log.Fatalln("Oops! Wasn't able to execute template", err)
	}

	//Accessing a map

	topTitles := map[string]string{"CEO": "Dev Ittycheria", "CTO": "Eliot Horowitz", "CFO": "Michael Gordon"}
	err = tpl.ExecuteTemplate(os.Stdout, "templatefileTitles.gohtml", topTitles)
	if err != nil {
		log.Fatalln("Oops! Wasn't able to execute template", err)
	}

	//Accessing a struct

	type products struct {
		Product  string //caps else it would be counted unexported
		Category string
	}

	p1 := products{Product: "Compass", Category: "Enterprise Tools"}
	err = tpl.ExecuteTemplate(os.Stdout, "templatefileStruct.gohtml", p1)
	if err != nil {
		log.Fatalln("Oops! Wasn't able to execute template", err)
	}

	//Accessing a list of products

	p2 := products{Product: "Charts", Category: "Enterprise Tools"}
	p3 := products{Product: "Database Server", Category: "Core Database"}
	p4 := products{Product: "Ops Manager", Category: "Enterprise Tools"}

	listOfProducts := []products{p1, p2, p3, p4}
	err = tpl.ExecuteTemplate(os.Stdout, "templatefileProducts.gohtml", listOfProducts)
	if err != nil {
		log.Fatalln("Oops! Wasn't able to execute template", err)
	}

	//Accessing structs of slices of structs

	type everything struct {
		Exec []string
		Prod []products
	}

	e1 := everything{Exec: listOfExecutives, Prod: listOfProducts}
	err = tpl.ExecuteTemplate(os.Stdout, "templatefileStructList.gohtml", e1)
	if err != nil {
		log.Fatalln("Oops! Wasn't able to execute template", err)
	}

	//Anonymous Templates - struct is created but not named
	data := struct {
		Exec []string
		Prod []products
	}{
		listOfExecutives,
		listOfProducts,
	}
	err = tpl.ExecuteTemplate(os.Stdout, "templatefileStructList.gohtml", data)
	if err != nil {
		log.Fatalln("Oops! Wasn't able to execute template", err)
	}
}
