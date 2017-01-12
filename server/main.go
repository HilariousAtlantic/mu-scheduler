package main

import "os"

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

	Section1 := Section{
		ID:       1,
		CourseID: 10,
		CRN:      "12345",
		Name:     "CSE 148",
	}
	Section2 := Section{
		ID:       2,
		CourseID: 20,
		CRN:      "54321",
		Name:     "CSE 283",
	}
	Section3 := Section{
		ID:       3,
		CourseID: 30,
		CRN:      "12345",
		Name:     "CSE 148",
	}
	Course1 := Course{
		ID:      10,
		TermID:  1,
		Subject: "CSE",
		Number:  "148",
		Title:   "Programming and stuff",
		Credits: "3.00",
	}
	Course2 := Course{
		ID:      20,
		TermID:  1,
		Subject: "CSE",
		Number:  "283",
		Title:   "Networking and stuff",
		Credits: "3.00",
	}
	Course3 := Course{
		ID:      30,
		TermID:  1,
		Subject: "CSE",
		Number:  "201",
		Title:   "Straight Bullshit",
		Credits: "3.00",
	}
	Meet1 := Meet{
		ID:         100,
		SectionID:  1,
		Days:       "MWF",
		StartTime:  "10:00",
		EndTime:    "12:00",
		Instructor: "SHIT ITS SOBEL",
		Location:   "BEN 101",
		StartDate:  "1/1/16",
		EndDate:    "1/1/20",
	}
	Meet2 := Meet{
		ID:         200,
		SectionID:  2,
		Days:       "MW",
		StartTime:  "8:00",
		EndTime:    "10:00",
		Instructor: "SHIT ITS SOBEL",
		Location:   "BEN 101",
		StartDate:  "1/1/16",
		EndDate:    "1/1/20",
	}
	Meet3 := Meet{
		ID:         300,
		SectionID:  3,
		Days:       "F",
		StartTime:  "14:00",
		EndTime:    "17:00",
		Instructor: "SHIT ITS SOBEL",
		Location:   "BEN 101",
		StartDate:  "1/1/16",
		EndDate:    "1/1/20",
	}
	Meet4 := Meet{
		ID:         400,
		SectionID:  3,
		Days:       "MTH",
		StartTime:  "10:30",
		EndTime:    "14:00",
		Instructor: "SHIT ITS SOBEL",
		Location:   "BEN 101",
		StartDate:  "1/1/16",
		EndDate:    "1/1/20",
	}
	Meet5 := Meet{
		ID:         500,
		SectionID:  2,
		Days:       "MF",
		StartTime:  "9:00",
		EndTime:    "12:00",
		Instructor: "SHIT ITS SOBEL",
		Location:   "BEN 101",
		StartDate:  "1/1/16",
		EndDate:    "1/1/20",
	}
	goodSchedules := make([][]Section, 0)

}
