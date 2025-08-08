package store

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=postgres user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("DB: Open %w", err)
	}

	fmt.Println("Connected to database...")
	return db, nil
}
