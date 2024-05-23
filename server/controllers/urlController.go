package controllers

import (
	"chisai/repositories"
	"chisai/services"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Url struct {
	Long_url string `json:"url"`
}

func HandleShortenRequest(c echo.Context) error {
	var body Url

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ERR: invalid request -> %s", err.Error()))
	}

	if body.Long_url == "" {
		return c.JSON(http.StatusBadRequest, errors.New("ERR: field [url] not provided").Error())
	}

	shortnedUrl := fmt.Sprintf("http://localhost:8080/%s", services.GenerateShortString(6))

	dbUrl, err := repositories.SaveURLs(body.Long_url, shortnedUrl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("ERR: failed to persist data -> %v", err))
	}

	return c.JSON(http.StatusOK, dbUrl)
}
