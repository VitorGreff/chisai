package main

import (
	"chisai/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, ip=${remote_ip}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/", func(c echo.Context) error { return controllers.GetURLs(c) })
	e.DELETE("/", func(c echo.Context) error { return controllers.ClearDatabase(c) })
	e.POST("/shorten", func(c echo.Context) error { return controllers.HandleShortenRequest(c) })
	e.GET("/r/:shortURL", func(c echo.Context) error { return controllers.HandleRedirectRequest(c) })

	e.Start(":8080")
}
