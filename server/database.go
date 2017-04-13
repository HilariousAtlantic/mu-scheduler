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

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
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
		end_date TEXT,
		attribute TEXT
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
		name TEXT NOT NULL,
		attribute TEXT NOT NULL
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
	INSERT INTO sections (course_id,crn,name,attribute)
	SELECT c.id,crn,s.name,attribute FROM staging s
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
	days, start_time, end_time, location, instructor,start_date,end_date,attribute)
	VALUES `

	insertGrades = `
	INSERT INTO grades (gpa, instructor, subject, number, year, season, section) VALUES`
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
	db *sqlx.DB
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

func (d *DBContext) open() *sqlx.DB {
	if d.db == nil {
		db, err := sql.Open("postgres", d.path())
		handleError(err)
		d.db = sqlx.NewDb(db, "postgres")
		d.db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	}
	return d.db
}

func createDatabase() {
	debug("Creating database...")
	output, err := exec.Command("createdb", "schedule_buddy", "-U", "schedule_buddy").CombinedOutput()
	if err != nil {
		debug(string(output))
		handleError(err)
	} else {
		debug("Database created")
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
	debug("Importing database...")

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
	debug("Importing grades...")

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
		debug(insertGrades)
	}

	debug("Grades Imported")
}

func getCoursesFromIDs(strIds string) (courses []*Course) {
	ids := make([]int, 0)
	for _, s := range strings.Split(strIds, ",") {
		id, _ := strconv.Atoi(s)
		ids = append(ids, id)
	}

	q, args, err := sqlx.In("SELECT * FROM courses WHERE id IN (?)", ids)
	handleError(err)

	db := dbContext.open()
	q = db.Rebind(q)
	err = db.Select(&courses, q, args...)
	handleError(err)

	return
}

func getMeetsFromSections(sections []*Section) (meets []*Meet) {
	ids := make([]int, len(sections))
	for i, section := range sections {
		ids[i] = section.ID
	}

	q, args, err := sqlx.In("SELECT * FROM meets WHERE section_id IN (?)", ids)
	handleError(err)

	db := dbContext.open()
	q = db.Rebind(q)
	err = db.Select(&meets, q, args...)
	handleError(err)

	return
}

func getSectionsFromCourses(courses []*Course) (sections []*Section) {
	ids := make([]int, len(courses))
	for i, course := range courses {
		ids[i] = course.ID
	}

	q, args, err := sqlx.In("SELECT * FROM sections WHERE course_id IN (?)", ids)
	handleError(err)

	db := dbContext.open()
	q = db.Rebind(q)
	err = db.Select(&sections, q, args...)
	handleError(err)

	return
}

func getCoursesFromDB(term string) (courses []*Course) {
	var err error
	db := dbContext.open()

	if term == "" {
		err = db.Select(&courses, "SELECT * FROM courses")
	} else {
		err = db.Select(&courses, "SELECT * FROM courses WHERE term_id = $1", term)
	}

	handleError(err)
	return
}

func getTermsFromDB() (terms []*Term) {
	db := dbContext.open()
	err := db.Select(&terms, "SELECT * FROM terms")
	handleError(err)
	return
}

func deleteDatabase() {
	debug("Deleting database...")

	output, err := exec.Command("dropdb", "schedule_buddy", "-U", "schedule_buddy").CombinedOutput()
	if err != nil {
		debug(err.Error(), string(output))
	} else {
		debug("Database deleted")
	}
}
