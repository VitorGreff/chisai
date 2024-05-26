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

func GetURLs(c echo.Context) error {
	urls, err := repositories.GetURLs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "ERR: unable to fetch URLs")
	}

	if len(urls) == 0 {
		return c.JSON(http.StatusOK, fmt.Sprintf("No URL registered"))
	}

	return c.JSON(http.StatusOK, urls)
}

func HandleShortenRequest(c echo.Context) error {
	var body Url

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ERR: invalid request -> %s", err.Error()))
	}

  if body.Long_url == ""{
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ERR: empty url"))
  }

	// check if URL is already on db
	existingUrl, err := repositories.GetURL(body.Long_url)
	if err == nil {
		return c.JSON(http.StatusOK, existingUrl)
	}

	shortnedUrl := fmt.Sprintf("http://localhost:8080/r/%s", services.GenerateShortString(4))

	newUrl, err := repositories.SaveURLs(body.Long_url, shortnedUrl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("ERR: failed to persist data -> %s", err.Error()))
	}

	return c.JSON(http.StatusOK, newUrl)
}

func ClearDatabase(c echo.Context) error {
	err := repositories.DeleleAllURLs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("ERR: failed to clear database -> %s", err.Error()))
	}

	return c.JSON(http.StatusOK, "Database cleaned")
}

func HandleRedirectRequest(c echo.Context) error {
	shortURL := c.Param("shortURL")
	if shortURL == "" {
		return c.JSON(http.StatusBadRequest, errors.New("ERR: missing path param [shortURL]").Error())
	}

	longURL, err := repositories.GetLongURL(fmt.Sprintf("http://localhost:8080/r/%s", shortURL))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("ERR: unable to redirect address").Error())
	}

	return c.Redirect(http.StatusFound, longURL)
}
