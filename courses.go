package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/buger/jsonparser"
)

func importCourses() {
	data, err := ioutil.ReadFile("courses.json")
	if err != nil {
		log.Fatal("Could not find courses.json")
	}

	// For each item in the top-level JSON array...
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		// Get the campus from each item
		campus, err := jsonparser.GetString(value, "campus")
		if err != nil {
			log.Fatal("Could not parse campus")
		}
		credits, err := jsonparser.GetString(value, "credits")
		if err != nil {
			log.Fatal("Could not parse credits")
		}
		id, err := jsonparser.GetString(value, "id")
		if err != nil {
			log.Fatal("Could not parse id")
		}
		number, err := jsonparser.GetString(value, "number")
		if err != nil {
			log.Fatal("Could not parse number")
		}
		section, err := jsonparser.GetString(value, "section")
		if err != nil {
			log.Fatal("Could not parse section")
		}
		subject, err := jsonparser.GetString(value, "subject")
		if err != nil {
			log.Fatal("Could not parse subject")
		}
		title, err := jsonparser.GetString(value, "title")
		if err != nil {
			log.Fatal("Could not parse title")
		}

		fmt.Println("campus: "+campus, "credits: "+credits, "id: "+id,
			" number: "+number, " section: "+section, " subject: "+subject, " title: "+title)


		//can get meets as a string by not as an array?
		fmt.Println(jsonparser.GetString(value, "meets"))

	})

	//this will successfully get all "meets" only, but they need to be associated with each course
	/*
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

		fmt.Println(jsonparser.GetString(value, "meets"))
	})
	*/
}
