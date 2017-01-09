package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "import":
			importCourses()
			importSemesters()
		case "createdb":
			createDatabase()
		case "dropdb":
			deleteDatabase()
		case "resetdb":
			deleteDatabase()
			createDatabase()
		default:
			startServer()
		}
	} else {
		startServer()
	}
}
