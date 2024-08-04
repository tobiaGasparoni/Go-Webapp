package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

// CreateDbConnection creates and returns the connection to the database
func CreateDbConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://citizix_user:S3cret@postgres:5432/citizix_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	log.Println("Connected to database!")

	return conn, nil
}
