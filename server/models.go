package main

type Course struct {
	ID       int    `json:"id"`
	Semester int    `json:"semester_id"`
	Name     string `json:"name"`
	Subject  string `json:"subject"`
	Number   string `json:"number"`
	Credits  string `json:"credits"`
}

type Semester struct {
	ID     int    `json:"id"`
	Season string `json:"Season"`
	Year   int    `json:"Year"`
	Name   string `json:"name"`
}

type Meet struct {
	ID         int    `json:"id"`
	Section_id string `json:"days"`
	Days       string `json:"days"`
	Start_time int    `json:"start_time"`
	End_time   int    `json:"end_time"`
	Instructor string `json:"instructor"`
	Location   string `json:"location"`
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
}

type Section struct {
	ID         int    `json:"id"`
	Section_id string `json:"section_id"`
	Section    string `json:"section"`
	Campus     string `json:"campus"`
}

func getMeetsFromJSON(parsed *parsedCoursesJSON) []*Meet {
	meets := make([]*Meet, 0)
	for _, course := range *parsed {
		for _, meet := range course.Meets {
			m := &Meet{
				ID:         -1,
				Section_id: course.ID,
				Days:       meet.Days,
				Start_time: -1,
				End_time:   -1,
				Instructor: meet.Instructor,
				Location:   meet.Location,
				Start_date: "Not set in import.go",
				End_date:   "Not set in import.go",
			}
			meets = append(meets, m)
		}
	}
	return meets
}

func getSectionsFromJSON(parsed *parsedCoursesJSON) []*Section {
	sections := make([]*Section, 0)
	for _, section := range *parsed {
		s := &Section{
			ID:         -1,
			Section_id: section.ID,
			Section:    section.Section,
			Campus:     section.Campus,
		}
		sections = append(sections, s)
	}
	return sections
}

func getCoursesFromJSON(parsed *parsedCoursesJSON) []*Course {
	courses := make([]*Course, 0)
	for _, course := range *parsed {
		c := &Course{
			ID:       -1,
			Semester: 3,
			Name:     course.Title,
			Subject:  course.Subject,
			Number:   course.Number,
			Credits:  course.Credits,
		}
		courses = append(courses, c)
	}
	return courses
}
