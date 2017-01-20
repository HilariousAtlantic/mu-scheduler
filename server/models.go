package main

type Term struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Course struct {
	Sections []Section `json:"sections"`
	ID       int       `json:"id"`
	TermID   int       `json:"term_id"`
	Subject  string    `json:"subject"`
	Number   string    `json:"number"`
	Title    string    `json:"title"`
	Credits  string    `json:"credits"`
}

type Section struct {
	Meets      []Meet  `json:"meets"`
	Tests      []Test  `json:"tests"`
	ID         int     `json:"id"`
	CourseID   int     `json:"course_id"`
	CRN        string  `json:"crn"`
	Name       string  `json:"name"`
	AverageGPA float64 `json:"average_grade"`
}

type Meet struct {
	ID         int    `json:"id"`
	SectionID  int    `json:"section_id"`
	Days       string `json:"days"`
	StartTime  int    `json:"start_time"`
	EndTime    int    `json:"end_time"`
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
	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
}

type Schedule struct {
	Sections []ScheduledCourse `json:"sections"`
}

type ScheduledCourse struct {
	CourseID  int `json:"course_id"`
	SectionID int `json:"section_id"`
}
