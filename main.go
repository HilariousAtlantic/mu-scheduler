package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "import":
			importCourses()
		default:
			runServer()
		}
	} else {
		runServer()
	}
}

func runServer() {
	e := echo.New()
	e.GET("/courses", courses)
	e.Logger.Fatal(e.Start(":8000"))
}

func courses(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	data, err := ioutil.ReadFile("courses.json")
	if err != nil {
		errMessage := `{"error": "Could not load courses"}`
		return c.String(http.StatusInternalServerError, errMessage)
	}

	data, err = ioutil.ReadFile("courses.json")

	return c.String(http.StatusOK, string(data))
}
