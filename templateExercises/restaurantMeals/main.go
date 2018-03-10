package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

// dish is one dish in the menu
type dish struct {
	Name  string
	Qty   string
	Price float64
	Desc  string
}

// meal refers to categorisation of a dish as Breakfast, Lunch or Dinner
type meal struct {
	Meal string
	Dish []dish
}

// menu will have a list of meals.
// each meal is a list of dishes
type menu []meal

// restaurant has a menu
type restaurant struct {
	Name string
	Menu menu
}

// restaurants has a list of restaurants in the city
type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./*.gohtml"))
}

func main() {
	r1 := restaurant{
		Name: "Kake Da Dhaba",
		Menu: menu{
			meal{
				Meal: "Breakfast",
				Dish: []dish{
					dish{
						Name:  "Dal Makhani",
						Qty:   "1",
						Price: 100,
						Desc:  "Pulse cooked with creamy layer",
					},
				},
			},
		},
	}
	fmt.Println(r1)

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", r1)
	if err != nil {
		log.Fatalln("Unable to execute template", err)
	}
}
