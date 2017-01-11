package main

type Term struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Course struct {
	ID      int    `json:"id"`
	TermID  int    `json:"term_id"`
	Subject string `json:"subject"`
	Number  string `json:"number"`
	Title   string `json:"title"`
	Credits string `json:"credits"`
}

type Section struct {
	ID       int    `json:"id"`
	CourseID int    `json:"course_id"`
	CRN      string `json:"crn"`
	Name     string `json:"name"`
}

type Meet struct {
	ID         int    `json:"id"`
	SectionID  int    `json:"section_id"`
	Days       string `json:"days"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Instructor string `json:"instructor"`
	Location   string `json:"location"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

type Test struct {
	ID        int    `json:"id"`
	SectionID int    `json:"section_id"`
	Date      string `json:"date"`
	Location  string `json:"location"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func getMeetsFromJSON(parsed *parsedCoursesJSON) []*Meet {
	meets := make([]*Meet, 0)
	sectionID := 0
	for _, course := range *parsed {
		for _, section := range course.Sections {
			for _, meet := range section.Meets {
				m := &Meet{
					ID:         -1,
					SectionID:  sectionID,
					Days:       meet.Days,
					StartTime:  meet.StartTime,
					EndTime:    meet.EndTime,
					Instructor: meet.Instructor,
					Location:   meet.Location,
					StartDate:  meet.StartDate,
					EndDate:    meet.EndDate,
				}
				meets = append(meets, m)
			}
			sectionID++
		}
	}
	return meets
}

func getTestsFromJSON(parsed *parsedCoursesJSON) []*Test {
	tests := make([]*Test, 0)
	sectionID := 0
	for _, course := range *parsed {
		for _, section := range course.Sections {
			for _, test := range section.Tests {
				t := &Test{
					ID:        -1,
					SectionID: sectionID,
					Date:      test.Date,
					Location:  test.Location,
					StartTime: test.StartTime,
					EndTime:   test.EndTime,
				}
				tests = append(tests, t)
			}
			sectionID++
		}
	}
	return tests
}

func getSectionsFromJSON(parsed *parsedCoursesJSON) []*Section {
	sections := make([]*Section, 0)
	courseID := 0
	for _, course := range *parsed {
		for _, section := range course.Sections {
			s := &Section{
				ID:       -1,
				CourseID: courseID,
				CRN:      section.CRN,
				Name:     section.Name,
			}
			sections = append(sections, s)
		}
		courseID++
	}
	return sections
}

func getCoursesFromJSON(parsed *parsedCoursesJSON) []*Course {
	courses := make([]*Course, 0)
	for _, course := range *parsed {
		c := &Course{
			ID:      -1,
			TermID:  course.TermID,
			Subject: course.Subject,
			Number:  course.Number,
			Title:   course.Title,
			Credits: course.Credits,
		}
		courses = append(courses, c)
	}
	return courses
}
