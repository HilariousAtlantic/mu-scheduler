package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	databasePath = "user=schedule_buddy dbname=schedule_buddy sslmode=disable"

	createStagingTable = `
	CREATE TABLE staging(
		type TEXT,
		term TEXT,
		crn TEXT,
		subject TEXT,
		number TEXT,
		title TEXT,
		name TEXT,
		credits TEXT,
		days TEXT,
		start_time INT,
		end_time INT,
		location TEXT,
		instructor TEXT,
		start_date TEXT,
		end_date TEXT
	);
	`

	createTermsTable = `
	CREATE TABLE terms (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	);
	`

	createCoursesTable = `
	CREATE TABLE courses (
		id SERIAL PRIMARY KEY,
		term_id INT NOT NULL REFERENCES terms,
		subject TEXT NOT NULL,
		number TEXT NOT NULL,
		title TEXT NOT NULL,
		credits TEXT NOT NULL
	);
	`

	createSectionsTable = `
	CREATE TABLE sections (
		id SERIAL PRIMARY KEY,
		course_id INT NOT NULL REFERENCES courses,
		crn TEXT NOT NULL,
		name TEXT NOT NULL
	);
	`

	createMeetsTable = `
	CREATE TABLE meets (
		id SERIAL PRIMARY KEY,
		section_id INT NOT NULL REFERENCES sections,
		days TEXT NOT NULL,
		start_time INT,
		end_time INT,
		location TEXT,
		instructor TEXT,
		start_date TEXT NOT NULL,
		end_date TEXT NOT NULL
	);
	`
	createTestsTable = `
	CREATE TABLE tests (
		id SERIAL PRIMARY KEY,
		section_id INT NOT NULL REFERENCES sections,
		date TEXT,
		start_time INT,
		end_time INT,
		location TEXT
	);
	`

	insertTerms = `
	INSERT INTO terms (name) SELECT DISTINCT term FROM staging;
	`

	insertCourses = `
	INSERT INTO courses (term_id,subject,number,title,credits)
	SELECT DISTINCT t.id,subject,number,title,credits FROM staging s
	JOIN terms t on t.name = s.term;
	`

	insertSections = `
	INSERT INTO sections (course_id,crn,name)
	SELECT c.id,crn,s.name FROM staging s
	JOIN terms t ON t.name = s.term
	JOIN courses c ON c.term_id = t.id AND c.subject = s.subject AND c.number = s.number
	WHERE type = 'section';
	`

	insertMeets = `
	INSERT INTO meets (section_id,days,start_time,end_time,location,instructor,start_date,end_date)
	SELECT sec.id,days,start_time,end_time,location,instructor,start_date,end_date FROM staging s
	JOIN sections sec ON sec.crn = s.crn
	WHERE type = 'meet' OR type = 'section';
				`
	insertTests = `
	INSERT INTO tests (section_id,start_time,end_time,location,date)
	SELECT sec.id,start_time,end_time,location,start_date FROM staging s
	JOIN sections sec ON sec.crn = s.crn
	WHERE type = 'test';
	`

	selectTests = `
	SELECT * FROM tests
	`
	selectMeets = `
	SELECT * FROM meets
	`
	selectSections = `
	SELECT * FROM sections
	`
	selectCourses = `
	SELECT * FROM courses
	`
	selectTerms = `
	SELECT * FROM terms
	`
)

var createTableStatements = [...]string{
	createStagingTable,
	createTermsTable,
	createCoursesTable,
	createSectionsTable,
	createMeetsTable,
	createTestsTable,
}

var insertStatements = [...]string{
	insertTerms,
	insertCourses,
	insertSections,
	insertMeets,
	insertTests,
}

type DBContext struct {
	db *sql.DB
}

var dbContext = new(DBContext)

func (d *DBContext) open() *sql.DB {
	if d.db == nil {
		var err error
		d.db, err = sql.Open("postgres", databasePath)
		if err != nil {
			handleError(err)
		}
	}
	return d.db
}

func createDatabase() {
	fmt.Println("Creating database...")
	output, err := exec.Command("createdb", "schedule_buddy", "-U", "schedule_buddy").CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		handleError(err)
	} else {
		fmt.Println("Database created")
	}

	db := dbContext.open()
	defer db.Close()

	for _, createTableStatement := range createTableStatements {
		_, err = db.Exec(createTableStatement)
		if err != nil {
			log.Fatalf("%q: %s\n", err, createTableStatement)
		}
	}
}

func importDatabase() {

	fmt.Println("Importing database...")

	f, err := os.Open("import/courses.csv")
	if err != nil {
		handleError(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	db := dbContext.open()
	defer db.Close()

	for _, line := range lines {
		var insertStaging = "INSERT INTO staging (type,term,crn,subject,number,title,name,credits,days,start_time,end_time,location,instructor,start_date,end_date) VALUES ("
		for _, value := range line {
			insertStaging += "'" + strings.Replace(value, "'", "''", -1) + "',"
		}
		insertStaging = insertStaging[0:len(insertStaging)-1] + ")"
		_, err = db.Exec(insertStaging)
		if err != nil {
			handleError(err)
			fmt.Println(insertStaging)
		}
	}

	for _, insertStatement := range insertStatements {
		_, err := db.Exec(insertStatement)
		if err != nil {
			handleError(err)
		}
	}

	fmt.Println("Database imported")
	fmt.Println("Testing functions")
	for _, course := range getCourseTree("100,230,231,680,1000,490,1800") {
		fmt.Print("section for Course ")
		fmt.Print(course.Subject)
		fmt.Print(course.Number)
		fmt.Print(" with ID: ")
		fmt.Println(course.ID)
		fmt.Println(course.Sections)
	}

	fmt.Println(doTimesOverlap(Section1, Section3))
	fmt.Println(doTimesOverlap(Section2, Section3))

}
func getCoursesFromIDString(ids string) []*Course {
	var courses []*Course
	db := dbContext.open()
	var rows *sql.Rows
	var err error
	where := "SELECT * FROM courses WHERE ID IN ("
	where += ids + ")"
	rows, err = db.Query(where)
	fmt.Println(where)
	defer rows.Close()
	for rows.Next() {
		course := &Course{}
		err = rows.Scan(&course.ID,
			&course.TermID,
			&course.Subject,
			&course.Number,
			&course.Title,
			&course.Credits)
		handleError(err)
		courses = append(courses, course)
	}
	err = rows.Err()
	handleError(err)
	return courses

}
func getMeetsFromSections(sections []*Section) []*Meet {
	var meets []*Meet
	db := dbContext.open()
	var rows *sql.Rows
	var err error
	where := "SELECT * FROM meets WHERE section_id IN ("
	for _, section := range sections {
		where += strconv.Itoa(section.ID) + ","
	}
	where = where[0 : len(where)-1]
	where += ")"
	rows, err = db.Query(where)
	defer rows.Close()
	for rows.Next() {
		meet := &Meet{}
		err = rows.Scan(&meet.ID,
			&meet.SectionID,
			&meet.Days,
			&meet.StartTime,
			&meet.EndTime,
			&meet.Location,
			&meet.Instructor,
			&meet.StartDate,
			&meet.EndDate)
		if err != nil {
			handleError(err)
		}
		meets = append(meets, meet)
	}
	err = rows.Err()
	if err != nil {
		handleError(err)
	}
	return meets
}
func getSectionsFromCourses(courses []*Course) []*Section {
	var sections []*Section
	db := dbContext.open()
	var rows *sql.Rows
	var err error
	where := "SELECT * FROM sections WHERE course_id IN ("
	for _, course := range courses {
		where += strconv.Itoa(course.ID) + ","
	}
	where = where[0 : len(where)-1]
	where += ")"
	rows, err = db.Query(where)

	if err != nil {
		handleError(err)
	}
	fmt.Println(rows)
	defer rows.Close()
	for rows.Next() {
		section := &Section{}
		err = rows.Scan(&section.ID,
			&section.CourseID,
			&section.CRN,
			&section.Name)
		if err != nil {
			handleError(err)
		}
		sections = append(sections, section)
	}
	err = rows.Err()
	if err != nil {
		handleError(err)
	}
	return sections
}

func getCoursesFromDB(term string) []*Course {
	var courses []*Course
	db := dbContext.open()
	var rows *sql.Rows
	var err error
	if term == "" {
		rows, err = db.Query(selectCourses)
		handleError(err)
	} else {
		rows, err = db.Query("SELECT * FROM courses WHERE term_id = $1", term)
		handleError(err)
	}
	defer rows.Close()
	for rows.Next() {
		course := &Course{}
		err = rows.Scan(&course.ID,
			&course.TermID,
			&course.Subject,
			&course.Number,
			&course.Title,
			&course.Credits)
		handleError(err)
		courses = append(courses, course)
	}
	err = rows.Err()
	handleError(err)
	return courses
}

func getSectionsFromDB() []*Section {
	var sections []*Section
	db := dbContext.open()
	rows, err := db.Query(selectSections)
	if err != nil {
		handleError(err)
	}
	defer rows.Close()
	for rows.Next() {
		section := &Section{}
		err = rows.Scan(&section.ID,
			&section.CourseID,
			&section.CRN,
			&section.Name)
		if err != nil {
			handleError(err)
		}
		sections = append(sections, section)
	}
	err = rows.Err()
	if err != nil {
		handleError(err)
	}
	return sections
}

func getTermsFromDB() []*Term {
	var terms []*Term
	db := dbContext.open()
	rows, err := db.Query(selectTerms)
	if err != nil {
		handleError(err)
	}
	defer rows.Close()
	for rows.Next() {
		term := &Term{}
		err = rows.Scan(
			&term.ID,
			&term.Name)
		if err != nil {
			handleError(err)
		}
		terms = append(terms, term)
	}
	err = rows.Err()
	if err != nil {
		handleError(err)
	}
	return terms
}

func getMeetsFromDB() []*Meet {
	var meets []*Meet
	db := dbContext.open()
	rows, err := db.Query(selectMeets)
	if err != nil {
		handleError(err)
	}
	defer rows.Close()
	for rows.Next() {
		meet := &Meet{}
		err = rows.Scan(&meet.ID,
			&meet.SectionID,
			&meet.Days,
			&meet.StartTime,
			&meet.EndTime,
			&meet.Instructor,
			&meet.Location,
			&meet.StartDate,
			&meet.EndDate)
		if err != nil {
			handleError(err)
		}
		meets = append(meets, meet)
	}
	err = rows.Err()
	if err != nil {
		handleError(err)
	}
	return meets
}

func getTestsFromDB() []*Test {
	var tests []*Test
	db := dbContext.open()
	rows, err := db.Query(selectTests)
	if err != nil {
		handleError(err)
	}
	defer rows.Close()
	for rows.Next() {
		test := &Test{}
		err = rows.Scan(&test.ID,
			&test.SectionID,
			&test.Date,
			&test.StartTime,
			&test.EndTime)
		if err != nil {
			handleError(err)
		}
		tests = append(tests, test)
	}
	err = rows.Err()
	if err != nil {
		handleError(err)
	}
	return tests
}

func deleteDatabase() {
	fmt.Println("Deleting database...")

	output, err := exec.Command("dropdb", "schedule_buddy", "-U", "schedule_buddy").CombinedOutput()
	if err != nil {
		fmt.Println(err.Error(), string(output))
	} else {
		fmt.Println("Database deleted")
	}
}
