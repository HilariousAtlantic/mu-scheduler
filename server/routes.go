package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func startServer() {
	e := echo.New()

	e.File("*", "dist/index.html")
	e.Static("app.js", "dist/app.js")

	api := e.Group("/api")

	api.GET("/terms", termsIndex)
	api.GET("/courses", coursesIndex)
	api.GET("/courses/:id", courseIndex)
	api.GET("/schedules", scheduleIndex)

	fmt.Println("Starting server on http://localhost:8000")

	e.Logger.Fatal(e.Start(":8000"))
}

func termsIndex(c echo.Context) error {
	terms := getTermsFromDB()
	return c.JSON(http.StatusOK, terms)
}

func coursesIndex(c echo.Context) error {
	term := c.QueryParam("term")
	courses := getCoursesFromDB(term)
	return c.JSON(http.StatusOK, courses)
}

func courseIndex(c echo.Context) error {
	id := c.Param("id")
	course := getCourseTree(id)
	return c.JSON(http.StatusOK, course)
}

func scheduleIndex(c echo.Context) error {
	courses := c.QueryParam("courses")
	schedules := findGoodSchedules(courses)
	return c.JSON(http.StatusOK, schedules)
}
