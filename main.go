package main

import (
	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes with Handler
	e.GET("/", index)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}

// Handler
func index(c echo.Context) error {
	return c.File("./index.html")
}
