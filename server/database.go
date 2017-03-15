package main

import (
	"bytes"
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
	createGradesTable = `
	CREATE TABLE grades (
		id SERIAL PRIMARY KEY,
		gpa NUMERIC NOT NULL,
		instructor TEXT NOT NULL,
		number TEXT NOT NULL,
		subject TEXT NOT NULL,
		year INT NOT NULL,
		season TEXT NOT NULL,
		section TEXT NOT NULL
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
	JOIN terms t ON t.name = s.term
	JOIN courses c ON c.term_id = t.id AND c.subject = s.subject AND c.number = s.number
	JOIN sections sec ON sec.crn = s.crn AND sec.course_id = c.id
	WHERE type = 'meet' OR type = 'section';
	`
	insertTests = `
	INSERT INTO tests (section_id,start_time,end_time,location,date)
	SELECT sec.id,start_time,end_time,location,start_date FROM staging s
	JOIN terms t ON t.name = s.term
	JOIN courses c ON c.term_id = t.id AND c.subject = s.subject AND c.number = s.number
	JOIN sections sec ON sec.crn = s.crn AND sec.course_id = c.id
	WHERE type = 'test';
	`

	insertStaging = `INSERT INTO staging
	(type, term, crn, subject, number, title, name, credits,
	days, start_time, end_time, location, instructor,start_date,end_date)
	VALUES `

	insertGrades = `
	INSERT INTO grades (gpa, instructor, subject, number, year, season, section) VALUES`

	selectCourses = `SELECT * FROM courses`

	selectTerms = `SELECT * FROM terms`
)

var createTableStatements = [...]string{
	createStagingTable,
	createTermsTable,
	createCoursesTable,
	createSectionsTable,
	createMeetsTable,
	createTestsTable,
	createGradesTable,
}

var insertStatements = [...]string{
	insertTerms,
	insertCourses,
	insertSections,
	insertMeets,
	insertTests,
}

var dbContext = new(DBContext)

type DBContext struct {
	db *sql.DB
}

func (d *DBContext) path() string {
	if !docker {
		return "user=schedule_buddy dbname=schedule_buddy sslmode=disable"
	}

	return fmt.Sprintf("user=%v dbname=%v password=%v host=postgres sslmode=disable",
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"))
}

func (d *DBContext) open() *sql.DB {
	if d.db == nil {
		var err error
		d.db, err = sql.Open("postgres", d.path())
		handleError(err)
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

	createTables()
}

func createTables() {
	db := dbContext.open()

	for _, createTableStatement := range createTableStatements {
		_, err := db.Exec(createTableStatement)
		if err != nil {
			log.Fatalf("%q: %s\n", err, createTableStatement)
		}
	}
}

func importDatabase() {
	fmt.Println("Importing database...")

	f, err := os.Open("import/courses.csv")
	handleError(err)
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	handleError(err)

	db := dbContext.open()

	debug("Building insert statement")

	insert := bytes.NewBufferString(insertStaging)
	for i, line := range lines {
		insert.WriteString("(")
		for j, value := range line {
			insert.WriteString(fmt.Sprintf(`'%v'`, strings.Replace(value, "'", "''", -1)))
			if j == len(line)-1 {
				insert.WriteString(")")
			} else {
				insert.WriteString(",")
			}
		}

		if i == len(lines)-1 {
			insert.WriteString(";")
		} else {
			insert.WriteString(",")
		}
	}

	debug("Inserting staging data")

	_, err = db.Exec(insert.String())
	if ok := handleError(err); !ok {
		debug(insert)
	}

	debug("Inserting data from staging")

	for _, insertStatement := range insertStatements {
		_, err := db.Exec(insertStatement)
		handleError(err)
	}

	debug("Database imported")
}

func importGradesDatabase() {
	fmt.Println("Importing grades...")

	f, err := os.Open("import/grades.csv")
	handleError(err)
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	handleError(err)

	db := dbContext.open()

	insert := bytes.NewBufferString(insertGrades)
	for i, line := range lines {
		insert.WriteString("(")
		for j, value := range line {
			insert.WriteString(fmt.Sprintf(`'%v'`, strings.Replace(value, "'", "''", -1)))
			if j == len(line)-1 {
				insert.WriteString(")")
			} else {
				insert.WriteString(",")
			}
		}

		if i == len(lines)-1 {
			insert.WriteString(";")
		} else {
			insert.WriteString(",")
		}
	}

	_, err = db.Exec(insert.String())
	if ok := handleError(err); !ok {
		fmt.Println(insertGrades)
	}

	fmt.Println("Grades Imported")
}

func getCoursesFromIDString(ids string) []*Course {
	var courses []*Course
	var rows *sql.Rows
	var err error

	where := "SELECT * FROM courses WHERE ID IN ("
	where += ids + ")"

	db := dbContext.open()
	rows, err = db.Query(where)

	debug(where)

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
	var rows *sql.Rows
	var err error

	db := dbContext.open()
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

	handleError(rows.Err())

	return meets
}

func getSectionsFromCourses(courses []*Course) []*Section {
	var sections []*Section
	var rows *sql.Rows
	var err error

	db := dbContext.open()
	where := "SELECT * FROM sections WHERE course_id IN ("

	for _, course := range courses {
		where += strconv.Itoa(course.ID) + ","
	}

	where = where[0 : len(where)-1]
	where += ")"
	rows, err = db.Query(where)
	handleError(err)

	debug(rows)

	defer rows.Close()
	for rows.Next() {
		section := &Section{}
		err = rows.Scan(&section.ID,
			&section.CourseID,
			&section.CRN,
			&section.Name)
		handleError(err)
		sections = append(sections, section)
	}

	handleError(rows.Err())

	return sections
}

func getCoursesFromDB(term string) []*Course {
	var courses []*Course
	var rows *sql.Rows
	var err error

	db := dbContext.open()

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

	handleError(rows.Err())

	return courses
}

func getTermsFromDB() []*Term {
	var terms []*Term

	db := dbContext.open()
	rows, err := db.Query(selectTerms)
	handleError(err)

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

	handleError(rows.Err())

	return terms
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
