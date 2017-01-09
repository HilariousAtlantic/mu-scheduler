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

	createSemestersTable = `
	CREATE TABLE semesters (
		id SERIAL PRIMARY KEY,
		season TEXT NOT NULL,
		year INT NOT NULL,
		name TEXT NOT NULL
	);
	`

	createCoursesTable = `
	CREATE TABLE courses (
		id SERIAL PRIMARY KEY,
		semester_id INT NOT NULL,
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
		start_time TEXT NOT NULL,
		end_time TEXT NOT NULL,
		instructor TEXT NOT NULL,
		location TEXT NOT NULL,
		start_date TEXT NOT NULL,
		end_date TEXT NOT NULL
	);
	`

	insertCourse = `
	INSERT INTO courses (semester_id, name, subject, number, credits) VALUES (?, ?, ?, ?, ?)
	`
	insertSemester = `
	INSERT INTO semesters (season, year, name) VALUES (?, ?, ?)
	`
	insertSection = `
	INSERT INTO sections (section,campus) VALUES (?, ?)
	`
	insertMeet = `
	INSERT INTO meets (section_id, days, start_time, end_time, instructor, location, start_date, end_date)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
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
	selectSemesters = `
	SELECT * FROM semesters
	`
)

var createTableStatements = [...]string{
	createCoursesTable,
	createSemestersTable,
	createSectionsTable,
	createMeetsTable,
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

func batchInsertSemesters(semesters []*Semester) {
	db := dbContext.open()
	tx, err := db.Begin()
	if err != nil {
		handleError(err)
	}
	stmt, err := tx.Prepare(pq.CopyIn("semesters", "season", "year", "name"))
	if err != nil {
		handleError(err)
	}
	defer stmt.Close()
	for _, semester := range semesters {
		_, err = stmt.Exec(semester.Season,
			semester.Year,
			semester.Name)
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
	stmt, err := tx.Prepare(pq.CopyIn("courses", "semester_id", "name", "subject", "number", "credits"))
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
		_, err = stmt.Exec(course.Semester,
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

func getSemestersFromDB() []*Semester {
	var semesters []*Semester
	db := dbContext.open()
	rows, err := db.Query(selectSemesters)
	if err != nil {
		handleError(err)
	}
	defer rows.Close()
	for rows.Next() {
		semester := &Semester{}
		err = rows.Scan(
			&semester.ID,
			&semester.Season,
			&semester.Year,
			&semester.Name)
		if err != nil {
			handleError(err)
		}
		semesters = append(semesters, semester)
	}
	err = rows.Err()
	if err != nil {
		handleError(err)
	}
	return semesters
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

func getCoursesFromDB() []*Course {
	var courses []*Course
	db := dbContext.open()
	rows, err := db.Query(selectCourses)
	if err != nil {
		handleError(err)
	}
	defer rows.Close()
	for rows.Next() {
		course := &Course{}
		err = rows.Scan(&course.ID,
			&course.Semester,
			&course.Name,
			&course.Subject,
			&course.Number,
			&course.Credits)
		if err != nil {
			handleError(err)
		}
		courses = append(courses, course)
	}
	err = rows.Err()
	if err != nil {
		handleError(err)
	}
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
