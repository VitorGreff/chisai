package main

import (
	"chisai/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/hello", func(c echo.Context) error { return c.JSON(http.StatusOK, "Hello, chisai!") })
	e.POST("/shorten", func(c echo.Context) error { return controllers.HandleShortenRequest(c)})

	e.Start(":8080")
}
