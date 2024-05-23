package controllers

import (
	"chisai/database"
	urlhandlers "chisai/urlHandlers"
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

type Url struct {
	Long_url string `json:"url"`
}

func ShortenURL(c echo.Context) (string, error) {
	var body Url

	if err := c.Bind(&body); err != nil  {
    return "", fmt.Errorf("Invalid URL provided. ERR: %v", err)
	}

  if body.Long_url == ""{
    return "", fmt.Errorf("Empty URL.")
  }

	shortnedUrl := fmt.Sprintf("http://localhost:8080/%s", urlhandlers.GenerateShortString(6))

	ctx := context.Background()

	conn, err := database.StartConnection(ctx)
	if err != nil {
		return "", err
	}

	query := database.New(conn)
	_, err = query.CreateLink(ctx, database.CreateLinkParams{
		LongUrl:  body.Long_url,
		ShortUrl: shortnedUrl,
	})
	if err != nil {
		return "", fmt.Errorf("ERR writing data on db.")
	}

	return shortnedUrl, nil
}
