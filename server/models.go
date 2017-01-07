package main

type Course struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Number  string `json:"number"`
	Credits string `json:"credits"`
}
type Meet struct {
	Days       string `json:"days"`
	Start_time int    `json:"start_time"`
	End_time   int    `json:"end_time"`
	Instructor string `json:"instructor"`
	Location   string `json:"location"`
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
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

//un-finished funciton, coming back later.
/*
func getMeetsFropmJSON(parsed *parsedCoursesJSON) []*Meet {
	meets := make([]*Meet, 0)
	for _, meet := range *parsed {
		m := &Meet{
			fmt.Printf(m)
		}

	}
}
*/
