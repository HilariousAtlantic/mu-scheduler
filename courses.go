package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

type parsedCoursesJSON []struct {
	Campus  string
	Credits string
	ID      string
	Meets   []struct {
		Date       string
		Days       string
		Instructor string
		Location   string
		Time       string
	}
	Number  string
	Section string
	Subject string
	Title   string
}

func importCourses() {
	data, err := ioutil.ReadFile("json/sections.json")
	if err != nil {
		log.Fatal("Could not find courses.json")
	}

	var parsed parsedCoursesJSON
	json.Unmarshal(data, &parsed)
	validateCourses(&parsed)
}

func validateCourses(parsed *parsedCoursesJSON) {
	creditsRegex, _ := regexp.Compile("\\d")
	numberRegex, _ := regexp.Compile("\\d+")
	campusRegex, _ := regexp.Compile("^[LVOHM]$")
	sectionRegex, _ := regexp.Compile("[A-Z\\d]")
	subjectRegex, _ := regexp.Compile("[A-Z]{3}")
	titleRegex, _ := regexp.Compile("[A-Za-z\\s]+")
	instructorRegex, _ := regexp.Compile("[a-zA-Z()]")
	locationRegex, _ := regexp.Compile("[A-Za-z\\d+]")
	dateRegex, _ := regexp.Compile("\\d+\\/\\d+-\\d+\\/\\d+")
	daysRegex, _ := regexp.Compile("[MWTRF]+")
	timeRegex, _ := regexp.Compile("(\\d+:\\d+ \\s?(am|pm)-\\d+:\\d+ \\s?(am|pm))")

	for _, data := range *parsed {
		if !campusRegex.MatchString(data.Campus) {
			fmt.Printf("Invalid campus: %v\n", data.Campus)
		}
		if !creditsRegex.MatchString(data.Credits) {
			fmt.Printf("Invalid credits: %v\n", data.Credits)
		}
		if !numberRegex.MatchString(data.Number) {
			fmt.Printf("Invalid number: %v\n", data.Number)
		}
		if !sectionRegex.MatchString(data.Section) {
			fmt.Printf("Invalid section: %v\n", data.Section)
		}
		if !subjectRegex.MatchString(data.Subject) {
			fmt.Printf("Invalid subject: %v\n", data.Subject)
		}
		if !titleRegex.MatchString(data.Subject) {
			fmt.Printf("Invalid title: %v\n", data.Title)
		}
		for _, meet := range data.Meets {
			if !dateRegex.MatchString(meet.Date) {
				fmt.Printf("Invalid date: %v\n", meet.Date)
			}
			if !daysRegex.MatchString(meet.Days) {
				fmt.Printf("Invalid days: %v\n", meet.Days)
			}
			if !instructorRegex.MatchString(meet.Instructor) {
				fmt.Printf("Invalid instructor: %v\n", meet.Instructor)
			}
			if !locationRegex.MatchString(meet.Location) {
				fmt.Printf("Invalid location: %v\n", meet.Location)
			}
			if !timeRegex.MatchString(meet.Time) {
				fmt.Printf("Invalid time: %v\n", meet.Time)
			}
		}
	}
}
