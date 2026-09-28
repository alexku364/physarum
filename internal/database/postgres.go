package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://localhost/physarum")
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	return conn, nil
}
