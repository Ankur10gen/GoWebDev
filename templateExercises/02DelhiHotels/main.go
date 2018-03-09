package main

import (
	"context"
	"encoding/json"
	"fmt"
	//"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./*.gohtml"))
}

type hotel struct {
	Name    string
	Address string
	City    string
	Pincode string
	Region  string
}

func main() {

	h1 := hotel{
		Name:    "Leela",
		Address: "Near Ambience Mall",
		City:    "Gurugram",
		Pincode: "122002",
		Region:  "North",
	}

	h2 := hotel{
		Name:    "Westin",
		Address: "Near IFFCO chowk",
		City:    "Gurugram",
		Pincode: "122001",
		Region:  "North",
	}

	hotels := []hotel{h1, h2}
	//fmt.Println(hotels)

	err := tpl.ExecuteTemplate(os.Stdout, "hotelstpl.gohtml", hotels)
	if err != nil {
		log.Fatalln("Something's not right with the template", err)
	}

	op, err := json.Marshal(h1)
	if err != nil {
		log.Fatalln("Wasn't able to marshal to json", err)
	}

	fmt.Println(string(op))

	enc := json.NewEncoder(os.Stdout)
	d := map[string]string{"a": "b"}
	enc.Encode(d)

	//Writing things to MongoDB

	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatalln("Unable to connect to database", err)
	}

	db := client.Database("hoteldb")
	coll := db.Collection("hotels")
	_, err = coll.InsertOne(context.Background(), h1)

	if err != nil {
		log.Fatalln("Unable to write to the collection", err)
	}

}
