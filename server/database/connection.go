package database

import (
	"context"
	"fmt"
	"os"

	"github.com/Valgard/godotenv"
	"github.com/jackc/pgx/v5"
)

func StartConnection(ctx context.Context) (*pgx.Conn, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, fmt.Errorf("Unable to get environment variables")
	}

	conn, err := pgx.Connect(ctx, fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DBNAME")))

	if err != nil {
		return nil, fmt.Errorf("Unable to establish connection.")
	}
	return conn, nil
}
