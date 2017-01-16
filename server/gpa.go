package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	sectionsURI     string = "http://grdist.miamioh.edu/php/getClasses.php?dept=%v&num=%v&inst=&from=%v&to=%v&iid=-1&did=%v&sem=&loc=O"
	departmentIDURI string = "http://grdist.miamioh.edu/php/getDID.php?dept=%s"
)

type departmentIDJSON struct {
	DepartmentID int `json:"did"`
}

type sectionsJSON struct {
	GPA           float32 `json:"avggpa"`
	Year          int     `json:"Year"`
	SubjectCode   string  `json:"NameShort"`
	CourseNumber  int     `json:"number"`
	ProfessorName string  `json:"name"`
}

type subjectMap map[string]uint

func (s *sectionsJSON) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n",
		s.GPA, s.Year, s.SubjectCode, s.CourseNumber, s.ProfessorName)
}

func importGPAs() {
	total := 0
	subjects := getSubjects()
	courses := getCoursesFromDB("")
	start := time.Now()
	count := 0
	for _, course := range courses {
		count++
		total += fetchGPAsForCourse(course, &subjects)
		fmt.Printf("%v/%v %v %v\n", count, len(courses), total, time.Since(start))
	}
}

func fetchGPAsForCourse(course *Course, subjects *subjectMap) int {
	//"http://grdist.miamioh.edu/php/getClasses.php?dept=%v&num=%v&inst=&from=%v&to=%v&iid=-1&did=%v&sem=&loc=O"
	res, err := http.Get(fmt.Sprintf(sectionsURI, course.Subject, course.Number, 2000, 2017, (*subjects)[course.Subject]))
	handleError(err)
	body, _ := ioutil.ReadAll(res.Body)

	var decoded []sectionsJSON
	json.Unmarshal(body, &decoded)
	return len(decoded)
}

func getSubjects() subjectMap {
	subjects := make(subjectMap)

	subjectCodes := getSubjectsFromDB()
	for _, subjectCode := range subjectCodes {
		res, err := http.Get(fmt.Sprintf(departmentIDURI, subjectCode))
		handleError(err)
		body, _ := ioutil.ReadAll(res.Body)

		var decoded []departmentIDJSON
		json.Unmarshal(body, &decoded)

		if len(decoded) > 0 && decoded[0].DepartmentID > -1 {
			subjects[subjectCode] = uint(decoded[0].DepartmentID)
		} else {
			handleError(fmt.Sprintf("Subject %v does not have a department ID", subjectCode))
		}
	}

	return subjects
}
