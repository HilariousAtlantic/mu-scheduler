package main

type Course struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Number  string `json:"number"`
	Credits string `json:"credits"`
}

func getCoursesFromJSON(parsed *parsedCoursesJSON) []*Course {
	courses := make([]*Course, 0)
	for _, course := range *parsed {
		c := &Course{
			ID:      -1,
			Name:    course.Title,
			Subject: course.Subject,
			Number:  course.Number,
			Credits: course.Credits,
		}
		courses = append(courses, c)
	}
	return courses
}
