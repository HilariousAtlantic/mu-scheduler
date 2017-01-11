package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func startServer() {
	e := echo.New()

	e.File("/", "index.html")
	e.File("/courses", "index.html")
	e.File("/schedules", "index.html")
	e.Static("/dist", "dist")

	api := e.Group("/api")

	api.GET("/courses", coursesIndex)
	api.GET("/terms", termsIndex)

	fmt.Println("Starting server on http://localhost:8000")

	e.Logger.Fatal(e.Start(":8000"))
}

func coursesIndex(c echo.Context) error {

	term := c.QueryParam("term")
	courses := getCoursesFromDB(term)
	return c.JSON(http.StatusOK, courses)

}

func termsIndex(c echo.Context) error {
	terms := getTermsFromDB()
	return c.JSON(http.StatusOK, terms)
}

func sectionsIndex(c echo.Context) error {
	sections := getSectionsFromDB()
	return c.JSON(http.StatusOK, sections)
}

func meetsIndex(c echo.Context) error {
	meets := getMeetsFromDB()
	return c.JSON(http.StatusOK, meets)
}
