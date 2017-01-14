package main

import (
	"os"
)

var GoodSchedules = make([][]Section, 0)
var Course1 = Course{
	ID:      10,
	TermID:  1,
	Subject: "CSE",
	Number:  "148",
	Title:   "Programming and stuff",
	Credits: "3.00",
}
var Course2 = Course{
	ID:      20,
	TermID:  1,
	Subject: "CSE",
	Number:  "283",
	Title:   "Networking and stuff",
	Credits: "3.00",
}
var Course3 = Course{
	ID:      30,
	TermID:  1,
	Subject: "CSE",
	Number:  "201",
	Title:   "Straight Bullshit",
	Credits: "3.00",
}
var Section1 = Section{
	Meets:    []Meet{Meet1},
	ID:       6,
	CourseID: 10,
	CRN:      "12345",
	Name:     "CSE 148",
}
var Section2 = Section{
	Meets:    []Meet{Meet2, Meet3},
	ID:       2992,
	CourseID: 20,
	CRN:      "54321",
	Name:     "CSE 283",
}
var Section3 = Section{
	Meets:    []Meet{Meet4},
	ID:       2,
	CourseID: 30,
	CRN:      "12345",
	Name:     "CSE 148",
}
var Meet1 = Meet{
	ID:         100,
	SectionID:  1,
	Days:       "MWF",
	StartTime:  610,
	EndTime:    730,
	Instructor: "SHIT ITS SOBEL",
	Location:   "BEN 101",
	StartDate:  "1/1/16",
	EndDate:    "1/1/20",
}
var Meet2 = Meet{
	ID:         200,
	SectionID:  2,
	Days:       "MW",
	StartTime:  480,
	EndTime:    600,
	Instructor: "SHIT ITS SOBEL",
	Location:   "BEN 101",
	StartDate:  "1/1/16",
	EndDate:    "1/1/20",
}
var Meet3 = Meet{
	ID:         300,
	SectionID:  3,
	Days:       "M",
	StartTime:  840,
	EndTime:    1020,
	Instructor: "SHIT ITS SOBEL",
	Location:   "BEN 101",
	StartDate:  "1/1/16",
	EndDate:    "1/1/20",
}
var Meet4 = Meet{
	ID:         400,
	SectionID:  3,
	Days:       "MTH",
	StartTime:  630,
	EndTime:    1020,
	Instructor: "SHIT ITS SOBEL",
	Location:   "BEN 101",
	StartDate:  "1/1/16",
	EndDate:    "1/1/20",
}
var Meet5 = Meet{
	ID:         500,
	SectionID:  2,
	Days:       "MF",
	StartTime:  510,
	EndTime:    720,
	Instructor: "SHIT ITS SOBEL",
	Location:   "BEN 101",
	StartDate:  "1/1/16",
	EndDate:    "1/1/20",
}
var Courses = []Course{Course1, Course2, Course3}
var Sections = []Section{Section1, Section2, Section3}
var Meets = []Meet{Meet1, Meet2, Meet3, Meet4, Meet5}

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "import":
			importDatabase()
		case "createdb":
			createDatabase()
		case "dropdb":
			deleteDatabase()
		case "resetdb":
			deleteDatabase()
			createDatabase()
		default:
			startServer()
		}
	} else {
		startServer()
	}
}
