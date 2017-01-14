package main

import (
	"fmt"
	"strings"
)

//without comments
//goodSchedules := make([[]Section]Course,0)
/*
func findGoodSchedules(courses *[]Course, selectedSections *[]Section) {

	if len(*selectedSections) == len(*courses) {
		GoodSchedules = append(GoodSchedules, *selectedSections)
		return
	}
SKIPCOURSE:
	for _, course := range *courses {

		for _, selectedSection := range *selectedSections {
			if selectedSection.CourseID == course.ID {
				continue SKIPCOURSE
			}
		}
		for _, section := range getSectionsFromCourse(*course) {
			for _, selectedSection := range *selectedSections {

				if !(section.startTime > selectedSection.startTime &&
					section.startTime < selectedSection.endTime ||
					selectedSection.startTime > section.startTime &&
						selectedSection.startTime < section.startTime) {
					*selectedSections = append(*selectedSections, section)
					findGoodSchedules(courses, selectedSections)
				}
			}
			return
		}
	}
}
*/
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
				fmt.Println(meetB)
				fmt.Println(meetA)
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
