package repositories

import (
	"chisai/database"
	"context"
)

func SaveURLs(long_url, short_url string) (database.Url, error) {
	ctx := context.Background()

	conn, err := database.StartConnection(ctx)
	if err != nil {
		return database.Url{}, err
	}

	query := database.New(conn)
	dbUrl, err := query.CreateLink(ctx, database.CreateLinkParams{
		LongUrl:  long_url,
		ShortUrl: short_url,
	})
	if err != nil {
		return database.Url{}, err
	}

	return dbUrl, nil
}
