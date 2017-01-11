package main

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"

	"github.com/lib/pq"
)

const (
	databasePath = "user=schedule_buddy dbname=schedule_buddy sslmode=disable"

	createTermsTable = `
	CREATE TABLE terms (
		id SERIAL PRIMARY KEY,
		season TEXT NOT NULL,
		year INT NOT NULL,
		name TEXT NOT NULL
	);
	`

	createCoursesTable = `
	CREATE TABLE courses (
		id SERIAL PRIMARY KEY,
		term_id INT NOT NULL,
		name TEXT NOT NULL,
		subject TEXT NOT NULL,
		number TEXT NOT NULL,
		credits TEXT NOT NULL
	);
	`

	createSectionsTable = `
	CREATE TABLE sections (
		id SERIAL PRIMARY KEY,
		section TEXT NOT NULL,
		campus TEXT NOT NULL
	);
	`

	createMeetsTable = `
	CREATE TABLE meets (
		id SERIAL PRIMARY KEY,
		section_id TEXT NOT NULL,
		days TEXT NOT NULL,
		start_time TEXT,
		end_time TEXT,
		instructor TEXT,
		location TEXT,
		start_date TEXT NOT NULL,
		end_date TEXT NOT NULL
	);
	`
	createTestsTable = `
	CREATE TABLE tests (
		id SERIAL PRIMARY KEY,
		section_id TEXT NOT NULL,
		date TEXT,
		location TEXT,
		start_time TEXT,
		end_time TEXT
	);
	`

	insertCourse = `
	INSERT INTO courses (term_id, name, subject, number, credits) VALUES (?, ?, ?, ?, ?)
	`
	insertTerm = `
	INSERT INTO terms (season, year, name) VALUES (?, ?, ?)
	`
	insertSection = `
	INSERT INTO sections (section,campus) VALUES (?, ?)
	`
	insertMeet = `
	INSERT INTO meets (section_id, days, start_time, end_time, instructor, location, start_date, end_date)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	insertTest = `
	INSERT INTO tests (section_id, days, date, location, start_time, end_time)
	
		VALUES (?, ?, ?, ?, ?)
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
	createCoursesTable,
	createTermsTable,
	createSectionsTable,
	createMeetsTable,
	createTestsTable,
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

func batchInsertTerms(terms []*Term) {
	db := dbContext.open()
	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	stmt, err := tx.Prepare(pq.CopyIn("terms", "season", "year", "name"))
	if err != nil {
		handleError(err)
	}
	defer stmt.Close()
	for _, term := range terms {
		_, err = stmt.Exec(term.Season,
			term.Year,
			term.Name)
		if err != nil {
			handleError(err)
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		handleError(err)
	}

	err = tx.Commit()
	if err != nil {
		handleError(err)
	}
}

func batchInsertCourses(courses []*Course) {
	existingCourses := map[Course]bool{}
	db := dbContext.open()
	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	stmt, err := tx.Prepare(pq.CopyIn("courses", "term_id", "name", "subject", "number", "credits"))
	if err != nil {
		handleError(err)
	}
	defer stmt.Close()
	for _, course := range courses {
		if existingCourses[*course] {
			continue
		} else {
			existingCourses[*course] = true
		}
		_, err = stmt.Exec(course.Term,
			course.Name,
			course.Subject,
			course.Number,
			course.Credits)
		if err != nil {
			handleError(err)
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		handleError(err)

	}
	err = tx.Commit()
	if err != nil {
		handleError(err)
	}
}

func batchInsertSections(sections []*Section) {
	db := dbContext.open()
	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	stmt, err := tx.Prepare(pq.CopyIn("sections", "section", "campus"))
	if err != nil {
		handleError(err)
	}
	defer stmt.Close()
	for _, section := range sections {
		_, err = stmt.Exec(section.Section, section.Campus)
		if err != nil {
			handleError(err)
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		handleError(err)
	}

	err = tx.Commit()
	if err != nil {
		handleError(err)
	}
}

func batchInsertMeets(meets []*Meet) {
	db := dbContext.open()
	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	stmt, err := tx.Prepare(pq.CopyIn("meets", "section_id", "days", "start_time", "end_time", "instructor",
		"location", "start_date", "end_date"))
	if err != nil {
		handleError(err)
	}
	defer stmt.Close()
	for _, meet := range meets {
		_, err = stmt.Exec(meet.Section_id,
			meet.Days,
			meet.Start_time,
			meet.End_time,
			meet.Instructor,
			meet.Location,
			meet.Start_date,
			meet.End_date)
		if err != nil {
			handleError(err)
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		handleError(err)

	}
	err = tx.Commit()
	if err != nil {
		handleError(err)
	}
}

func filterCourses(filters map[string][]string) {
	var where, orderBy string = "", ""

	//check for the filter "days" (what days do you not want?
	days, err := filters["days"]
	if err == true {
		for _, day := range days {
			if where == "" {
				where += (" meets.days LIKE " + day)
			} else {
				where += (" AND NOT meets.days LIKE " + day)
			}
		}
	}
	//check for the filter "timesAfter"
	times, err := filters["timesAfter"]
	if err == true {
		for _, time := range times {
			fmt.Println(time)
			//do time filtering stuff and add to where
		}
	}

	//include another filter for  ORDERBY gpa

	filter := "SELECT id, name, subject, number, credits FROM courses ORDER BY" + orderBy + "WHERE" + where
	db := dbContext.open()
	rows, err1 := db.Query(filter)
	if err1 != nil {
		handleError(err1)
	}
	fmt.Println(rows)
	//finish query, etc
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
			&section.Section,
			&section.Campus)
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
			&term.Season,
			&term.Year,
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
			&meet.Section_id,
			&meet.Days,
			&meet.Start_time,
			&meet.End_time,
			&meet.Instructor,
			&meet.Location,
			&meet.Start_date,
			&meet.End_date)
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

func getCoursesFromDB(term string) []*Course {
	var courses []*Course
	db := dbContext.open()
	var rows *sql.Rows
	var err error
	if term == "" {
		rows, err = db.Query(selectCourses)
		handleError(err)
	} else {
		rows, err = db.Query("SELECT * FROM courses WHERE term_id=?", term)
		handleError(err)
	}
	defer rows.Close()
	for rows.Next() {
		course := &Course{}
		err = rows.Scan(&course.ID,
			&course.Term,
			&course.Name,
			&course.Subject,
			&course.Number,
			&course.Credits)
		handleError(err)
		courses = append(courses, course)
	}
	err = rows.Err()
	handleError(err)
	return courses
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
