package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	sectionsURI     string = "http://grdist.miamioh.edu/php/getClasses.php?dept=%v&num=&inst=&from=%v&to=%v&iid=-1&did=%v&sem=&loc=O"
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

func (s *sectionsJSON) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n",
		s.GPA, s.Year, s.SubjectCode, s.CourseNumber, s.ProfessorName)
}

func importGPAs() {
	subjects := getSubjects()

	for code, id := range subjects {
		fetchGPAsFromSubject(code, id, 2016)
	}
}

func fetchGPAsFromSubject(subjectCode string, subjectID uint, year int) {
	res, err := http.Get(fmt.Sprintf(sectionsURI, subjectCode, year-1, year, subjectID))
	handleError(err)
	body, _ := ioutil.ReadAll(res.Body)

	var decoded []sectionsJSON
	json.Unmarshal(body, &decoded)
}

func getSubjects() map[string]uint {
	subjects := make(map[string]uint)

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
