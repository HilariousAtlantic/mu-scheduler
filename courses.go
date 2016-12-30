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

	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		campus, err := jsonparser.GetString(value, "campus")
		if err != nil {
			log.Fatal("Could not parse campus")
		}
		fmt.Println(campus)
	})
}
