package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

var fm = template.FuncMap{
	"fDateMDY":    timeConversion,
	"funnyString": funnyString,
}

func timeConversion(t time.Time) string {
	return t.Format("01-02-2006")
}

func funnyString(s string) string {
	return s + " HA HA!"
}

func main() {
	t := time.Now()

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", t)
	if err != nil {
		log.Fatalln("Couldn't execute template", err)
	}
}
