package main

import (
	"chisai/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/hello", func(c echo.Context) error { return c.JSON(http.StatusOK, "Hello, chisai!") })
	e.POST("/chisai", func(c echo.Context) error {
		url, err := controllers.ShortenURL(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, url)
	})

	e.Start(":8080")
}
