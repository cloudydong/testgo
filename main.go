package main

import (
	"github.com/cloudydong/testgo/controller"
	"github.com/cloudydong/testgo/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDB()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes with Handler
	e.GET("/", index)
	e.GET("corp", controller.GetCorp)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func index(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello world")
	return c.File("index.html")
}
