package main

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
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

	return c.String(http.StatusOK, string(data))
}
