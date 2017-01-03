package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	databasePath = "./db.sqlite3"
)

const (
	createCoursesTableStatement = `
	CREATE TABLE courses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		subject TEXT NOT NULL,
		number INTEGER NOT NULL,
		credits TEXT NOT NULL
	);
	`
	insertCourseStatement = `
	INSERT INTO courses(name, subject, number, credits) values(?, ?, ?, ?)
	`

	selectCourses = `
	SELECT id, name, subject, number, credits FROM courses
	`
)

type DB struct {
	db *sql.DB
}

var dbContext = new(DB)

func (d *DB) open() *sql.DB {
	if d.db == nil {
		var err error
		d.db, err = sql.Open("sqlite3", databasePath)
		if err != nil {
			log.Fatal(err)
		}
	}
	return d.db
}

func createDatabase() {
	db := dbContext.open()
	defer db.Close()

	_, err := db.Exec(createCoursesTableStatement)
	if err != nil {
		log.Fatalf("%q: %s\n", err, createCoursesTableStatement)
	}
}

func batchInsertCourses(courses []*Course) {
	existingCourseNames := map[string]bool{}
	db := dbContext.open()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(insertCourseStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for _, course := range courses {
		if existingCourseNames[course.Name] {
			continue
		} else {
			existingCourseNames[course.Name] = true
		}
		_, err = stmt.Exec(course.Name, course.Subject, course.Number, course.Credits)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
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
	os.Remove(databasePath)
}
