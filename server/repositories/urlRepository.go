package repositories

import (
	"chisai/database"
	"context"
)

func GetURL(long_url string) (database.Url, error) {
	ctx := context.Background()

	conn, err := database.StartConnection(ctx)
	if err != nil {
		return database.Url{}, err
	}
	defer conn.Close(ctx)

	query := database.New(conn)
	existingUrl, err := query.GetLink(ctx, long_url)
	if err != nil {
		return database.Url{}, err
	}

	return existingUrl, nil
}

func GetLongURL(short_url string) (string, error) {
	ctx := context.Background()

	conn, err := database.StartConnection(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Close(ctx)

	query := database.New(conn)
	long_url, err := query.GetLongURL(ctx, short_url)
	if err != nil {
		return "", err
	}

	return long_url, nil
}

func GetURLs() ([]database.Url, error) {
	ctx := context.Background()

	conn, err := database.StartConnection(ctx)
	if err != nil {
		return []database.Url{}, err
	}
	defer conn.Close(ctx)

	query := database.New(conn)
	urls, err := query.ListLinks(ctx)
	if err != nil {
		return []database.Url{}, err
	}

	return urls, nil
}

func SaveURLs(long_url, short_url string) (database.Url, error) {
	ctx := context.Background()

	conn, err := database.StartConnection(ctx)
	if err != nil {
		return database.Url{}, err
	}
	defer conn.Close(ctx)

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

func DeleleAllURLs() error {
	ctx := context.Background()

	conn, err := database.StartConnection(ctx)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	query := database.New(conn)
	err = query.DeleteAllLinks(ctx)
	if err != nil {
		return err
	}

	return nil
}
