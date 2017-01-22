package main

import (
	"database/sql"
	"fmt"
	"math"
	"strings"
)

// a helper function that is used to create the necessary
// variables to call findGoodSchedulesRecursive
func findGoodSchedules(ids string) []Schedule {

	goodSchedules := make([]Schedule, 0)

	//courses cannot be a pointer due to recursion, so convert
	coursesPointer := getCourseTree(ids)
	courses := make([]Course, 0)
	for _, coursePointer := range coursesPointer {
		courses = append(courses, *coursePointer)
	}

	selectedSections := make([]Section, 0)
	findGoodSchedulesRecursive(courses, selectedSections, &goodSchedules)
	//fmt.Println("good schedules")

	return goodSchedules

}

/*recursive function to find all good schedules.
  works by checking every combination until it finds an invalid time, then returning and checking
  the next section.
*/
func findGoodSchedulesRecursive(courses []Course, selectedSections []Section, goodSchedules *[]Schedule) {
	//base case; we found a good schedule. Append it and return.
	if len(selectedSections) == len(courses) {
		//fmt.Println("len is: " + string(len(courses)) + " sections are: ")
		//fmt.Println(selectedSections)
		var sections []ScheduledCourse
		goodScheduleAvgGPA := 0.0
		for _, selectedSection := range selectedSections {
			if !math.IsNaN(selectedSection.AverageGPA) {
				goodScheduleAvgGPA += selectedSection.AverageGPA
			}
			var scheduledCourse = ScheduledCourse{
				CourseID:  selectedSection.CourseID,
				SectionID: selectedSection.ID,
			}
			sections = append(sections, scheduledCourse)
		}
		var goodSchedule = Schedule{
			Sections:   sections,
			AverageGPA: goodScheduleAvgGPA,
		}

		//fmt.Printf("sched gpa is: %v", goodScheduleAvgGPA)
		//fmt.Println("good schedule added:")
		//fmt.Println(goodSchedule)
		*goodSchedules = append(*goodSchedules, goodSchedule)
		return
	}
	//skip a course and continue
SKIPCOURSE:
	for _, course := range courses {
		//fmt.Println("new course")
		for _, selectedSection := range selectedSections {
			//if we have a course in the selectedSections, we dont check the other sections of said course
			//fmt.Printf("courseID is: %v course.ID is: %v", selectedSection.CourseID, course.ID)
			if selectedSection.CourseID == course.ID {
				continue SKIPCOURSE
			}
		}
		//fmt.Println("approved")
	SKIPSECTION:
		for _, section := range course.Sections {
			//go through all selectedSections and make sure none overlap
			for _, selectedSection := range selectedSections {

				//if overlap, skip that section
				if doTimesOverlap(selectedSection, section) {
					continue SKIPSECTION
				}
			}
			//if none overlap, section is good, add to selectedSections
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

//returns a built out "tree" of courses.
//a tree means that the courses have sections and the sections have meets
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
	setCoursesGPA(courses)
	return courses
}
func setCoursesGPA(courses []*Course) {
	for _, course := range courses {
		for _, section := range course.Sections {
			meet := section.Meets[0]
			instructor := meet.Instructor
			index := strings.Index(instructor, ";")

			//there is not two professors
			if index != -1 {
				instructors := strings.Split(instructor, ";")
				//instructor1 := instructors[0]
				//instructors2 := instructors[1]
				avgGPA := (getAvgGPA(instructors[0], *course) + getAvgGPA(instructors[1], *course)) / 2
				section.AverageGPA = avgGPA
			} else {
				fmt.Println(getAvgGPA(instructor, *course))
				section.AverageGPA = getAvgGPA(instructor, *course)
			}
		}
	}
}

func getAvgGPA(instructor string, course Course) float64 {
	var avgGPA = 0.0
	var divideBy = 0.0
	instructor = strings.Replace(instructor, "(P)", "", -1)
	instructor = strings.TrimSpace(instructor)
	fmt.Printf("instructor after:%v.\n", instructor)
	db := dbContext.open()
	var rows *sql.Rows
	var err error
	query := "SELECT gpa FROM grades WHERE LOWER(instructor) LIKE LOWER('%" + instructor +
		"%') AND subject = '" + course.Subject +
		"' AND number = '" + course.Number + "';"
	//query := "SELECT gpa FROM grades WHERE LOWER(instructor) LIKE LOWER('%zmuda%')"
	rows, err = db.Query(query)
	handleError(err)
	defer rows.Close()
	for rows.Next() {
		divideBy++
		var gpa float64
		err = rows.Scan(&gpa)
		//fmt.Printf("in loop gpa: %v", gpa)
		avgGPA += gpa
		handleError(err)
	}
	err = rows.Err()
	handleError(err)
	avgGPA = (avgGPA / divideBy)
	avgGPA = float64(int(avgGPA*100)) / 100
	//fmt.Printf("avg gpa is: %v", avgGPA)
	return avgGPA

}

//checks if two sections have any meet times that over lap,
// ignoring meets that have no overlapping days
func doTimesOverlap(a, b Section) bool {
	for _, meetA := range a.Meets {
		for _, meetB := range b.Meets {
			//fmt.Printf("comparing meet: %v to meet: %v", meetA, meetB)
			if !containsSameDay(meetA.Days, meetB.Days) {
				//fmt.Println("different day")
				continue
			} else if meetA.StartTime <= meetB.EndTime &&
				meetB.StartTime <= meetA.EndTime {
				//fmt.Println(meetB)
				//fmt.Println(meetA)
				//fmt.Println("no overlap")
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
