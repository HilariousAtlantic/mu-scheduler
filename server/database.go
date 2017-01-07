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

	createCoursesTable = `
	CREATE TABLE courses (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		subject TEXT NOT NULL,
		number TEXT NOT NULL,
		credits TEXT NOT NULL
	);
	`

	createSemestersTable = `
	CREATE TABLE semesters (
		id SERIAL PRIMARY KEY,
		season TEXT NOT NULL,
		year INT NOT NULL,
		name TEXT NOT NULL
	);
	`

	createMeetsTable = `
	CREATE TABLE semesters (
		days TEXT NOT NULL,
		start_time INT NOT NULL,
		end_time INT NOT NULL,
		instructor TEXT NOT NULL,
		location TEXT NOT NULL,
		start_date TEXT NOT NULL,
		end_date TEXT NOT NULL
	);
	`

	insertCourse = `
	INSERT INTO courses (name, subject, number, credits) VALUES (?, ?, ?, ?)
	`
	insertSemester = `
	INSERT INTO semesterss (season, year, name) VALUES (?, ?, ?, ?)
	`

	selectCourses = `
	SELECT id, name, subject, number, credits FROM courses
	`
	selectSemesters = `
	SELECT id, season, year, name FROM semesters
	`
)

type DB struct {
	db *sql.DB
}

var dbContext = new(DB)

func (d *DB) open() *sql.DB {
	if d.db == nil {
		var err error
		d.db, err = sql.Open("postgres", databasePath)
		if err != nil {
			log.Fatal(err)
		}
	}
	return d.db
}

func createDatabase() {
	fmt.Println("Creating database...")

	output, err := exec.Command("createdb", "schedule_buddy", "-U", "schedule_buddy").CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		log.Fatal(err)
	} else {
		fmt.Println("Database created")
	}

	db := dbContext.open()
	defer db.Close()

	_, err = db.Exec(createCoursesTable)

	if err != nil {
		log.Fatalf("%q: %s\n", err, createCoursesTable)
	}
	_, err = db.Exec(createSemestersTable)

	if err != nil {
		log.Fatalf("%q: %s\n", err, createSemestersTable)
	}

}

func batchInsertSemesters(semesters []*Semester) {
	db := dbContext.open()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(pq.CopyIn("semesters", "season", "year", "name"))
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for _, semester := range semesters {
		_, err = stmt.Exec(semester.Season, semester.Year, semester.Name)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
func batchInsertCourses(courses []*Course) {
	//fmt.Println(courses)
	existingCourses := map[Course]bool{}
	db := dbContext.open()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(pq.CopyIn("courses", "name", "subject", "number", "credits"))
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for _, course := range courses {
		if existingCourses[*course] {
			continue
		} else {
			existingCourses[*course] = true
		}
		_, err = stmt.Exec(course.Name, course.Subject, course.Number, course.Credits)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)

	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

//needs finishing
func batchInsertMeets(meets []*Meet) {
	existingCourses := map[Course]bool{}
	db := dbContext.open()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(pq.CopyIn("courses", "name", "subject", "number", "credits"))
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for _, course := range courses {
		if existingCourses[*course] {
			continue
		} else {
			existingCourses[*course] = true
		}
		_, err = stmt.Exec(course.Name, course.Subject, course.Number, course.Credits)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)

	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

//UNFINISHED function for filtering
//map["filter name"] []list of things to be filtered
//ex: map: "days" -> ["M", "W"]
//	   "timesAfter" -> ["9:00"]
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
		log.Fatal(err1)
	}
	fmt.Println(rows)
	//finish query, etc
}

func getCoursesFromDB() []*Course {
	var courses []*Course
	db := dbContext.open()
	rows, err := db.Query(selectCourses)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		course := &Course{}
		err = rows.Scan(&course.ID, &course.Name, &course.Subject, &course.Number, &course.Credits)
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, course)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
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
