package main

import "os"

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "import":
			importCourses()
			importTerms()
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
