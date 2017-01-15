package main

import (
	"fmt"
	"strings"
)

//without comments
//goodSchedules := make([[]Section]Course,0)

func findGoodSchedules(ids string) []Schedule {

	goodSchedules := make([]Schedule, 0)
	coursesPointer := getCourseTree(ids)
	courses := make([]Course, 0)
	for _, coursePointer := range coursesPointer {
		courses = append(courses, *coursePointer)
	}

	selectedSections := make([]Section, 0)
	findGoodSchedulesRecursive(courses, selectedSections, &goodSchedules)
	fmt.Println("good schedules")
	for _, goodSchedule := range goodSchedules {
		goodSchedule.Courses = courses
		fmt.Println(goodSchedule.Sections)
	}
	return goodSchedules

}

func findGoodSchedulesRecursive(courses []Course, selectedSections []Section, goodSchedules *[]Schedule) {

	if len(selectedSections) == len(courses) {
		//fmt.Println("len is: " + string(len(courses)) + " sections are: ")
		//fmt.Println(selectedSections)
		var newGoodSchedule Schedule
		newGoodSchedule.Sections = selectedSections
		*goodSchedules = append(*goodSchedules, newGoodSchedule)
		return
	}
SKIPCOURSE:
	for _, course := range courses {
		//fmt.Println("new course")
		for _, selectedSection := range selectedSections {
			//fmt.Printf("courseID is: %v course.ID is: %v", selectedSection.CourseID, course.ID)
			if selectedSection.CourseID == course.ID {
				continue SKIPCOURSE
			}
		}
		//fmt.Println("approved")
	SKIPSECTION:
		for _, section := range course.Sections {
			for _, selectedSection := range selectedSections {

				if doTimesOverlap(selectedSection, section) {
					continue SKIPSECTION
				}
			}
			//fmt.Println(selectedSections)
			selectedSections = append(selectedSections, section)
			//fmt.Println(selectedSections)
			findGoodSchedulesRecursive(courses, selectedSections, goodSchedules)
			selectedSections = selectedSections[:len(selectedSections)-1]
		}
		//fmt.Println("none worked")
		return
	}
	return
}

func getCourseTree(ids string) []*Course {

	courses := getCoursesFromIDString(ids)
	sections := getSectionsFromCourses(courses)
	meets := getMeetsFromSections(sections)

	for _, section := range sections {
		if len(meets) == 0 {
			break
		}
		for _, meet := range meets {
			if meet.SectionID == section.ID {
				section.Meets = append(section.Meets, *meet)
			}
		}
	}
	for _, course := range courses {
		for _, section := range sections {
			if section.CourseID == course.ID {
				course.Sections = append(course.Sections, *section)
			}
		}
	}
	return courses
}

//checks if two sections have any meet times that over lap,
// ignoring meets that have no overlapping days
func doTimesOverlap(a, b Section) bool {
	for _, meetA := range a.Meets {
		for _, meetB := range b.Meets {
			if !containsSameDay(meetA.Days, meetB.Days) {
				break
			} else if meetA.StartTime <= meetB.EndTime &&
				meetB.StartTime <= meetA.EndTime {
				//fmt.Println(meetB)
				//fmt.Println(meetA)
				return true
			}
		}
	}
	return false
}

//returns true if the two strings share a similar character
//accounts for TBA and any upper/lower case issues
func containsSameDay(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	if a == "tba" || b == "tba" {
		return false
	}
	m := make(map[string]bool)
	for _, c := range a {
		s := string(c)
		m[s] = true
	}
	for _, c := range b {
		s := string(c)
		val, _ := m[s]
		if val == true {
			return true
		}
	}
	return false
}
