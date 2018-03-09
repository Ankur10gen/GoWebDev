package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	return s[:3]
	//return s
}

func main() {
	sage1 := sage{
		Name:  "Mahatma Gandhi",
		Motto: "Be the change you want to see in the world",
	}

	sage2 := sage{
		Name:  "Rob Pike",
		Motto: "Concurrency is not Parallelism",
	}

	sages := []sage{sage1, sage2}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln("Whoops! Wasn't able to execute the template", err)
	}
}
