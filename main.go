package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	// If there are 2 or more command line
	// arguments then check what the second
	// one is.
	// i.e. go run *.go <os.Args[1]>
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "import":
			importCourses()
		case "createdb":
			createDatabase()
		case "rmdb":
			deleteDatabase()
		default:
			runServer()
		}
	} else {
		runServer()
	}
}

func runServer() {
	e := echo.New()
	e.GET("/courses", coursesIndex)
	e.File("/", "index.html")
	e.Static("/dist", "dist")
	e.Logger.Fatal(e.Start(":8000"))
}

func coursesIndex(c echo.Context) error {
	courses := getCoursesFromDB()
	return c.JSON(http.StatusOK, courses)
}
