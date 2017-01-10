package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
)

type parsedCoursesJSON []struct {
	Subject  string
	Number   string
	Title    string
	Credits  string
	Sections []struct {
		CRN    int
		name   string
		title  string
		Campus string
		tests  []struct {
		}
		Meets []struct {
			Days       string
			Start_time string
			End_time   string
			Instructor string
			Start_date string
			End_date   string
			Location   string
		}
	}
}

func importCourses() {
	data, err := ioutil.ReadFile("import/json/spring-2017.json")
	if err != nil {
		handleError("Could not find sections.json")
	}

	var parsed parsedCoursesJSON
	json.Unmarshal(data, &parsed)
	//printInvalidCourses(&parsed)
	courses := getCoursesFromJSON(&parsed)
	sections := getSectionsFromJSON(&parsed)
	meets := getMeetsFromJSON(&parsed)
	batchInsertCourses(courses)
	batchInsertSections(sections)
	batchInsertMeets(meets)
}

func importSemesters() {
	semesters := []*Semester{
		&Semester{-1, "Fall", 2016, "Fall 2016"},
		&Semester{-1, "Winter", 2016, "Winter 2016"},
		&Semester{-1, "Spring", 2017, "Spring 2017"}}
	batchInsertSemesters(semesters)
}

//was used for parsing bad values, but its now done on front-end. Saving for future.
//TODO: move this to a utils.go file for future
/*
func printInvalidCourses(parsed *parsedCoursesJSON) {
	creditsRegex, _ := regexp.Compile("[\\d.-]")
	numberRegex, _ := regexp.Compile("\\d+([A-Z])?$")
	campusRegex, _ := regexp.Compile("^[LVOHM]$")
	sectionRegex, _ := regexp.Compile("[A-Z\\d]")
	subjectRegex, _ := regexp.Compile("[A-Z]{3}")
	titleRegex, _ := regexp.Compile("[A-Za-z\\s]+")
	instructorRegex, _ := regexp.Compile("[a-zA-Z()]")
	locationRegex, _ := regexp.Compile("[A-Za-z\\d+]")
	dateRegex, _ := regexp.Compile("\\d+\\/\\d+-\\d+\\/\\d+")
	daysRegex, _ := regexp.Compile("([MWTRFSU])+")
	timeRegex, _ := regexp.Compile("(\\d+:\\d+ \\s?(am|pm)-\\d+:\\d+ \\s?(am|pm))|TBA")

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
			if meet.Days != "" && !daysRegex.MatchString(meet.Days) {
				fmt.Printf("Invalid days: %v\n", meet.Days)
			}
			if !instructorRegex.MatchString(meet.Instructor) {
				fmt.Printf("Invalid instructor: %v\n", meet.Instructor)
			}
			if !locationRegex.MatchString(meet.Location) {
				fmt.Printf("Invalid location: %v\n", meet.Location)
			}
			if meet.Time != "" && !timeRegex.MatchString(meet.Time) {
				fmt.Printf("Invalid time: %v\n", meet.Time)
			}
		}
	}
}*/
