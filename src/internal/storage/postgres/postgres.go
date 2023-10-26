package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	cfg "github.com/sabrs0/bmstu-web/src/internal/config/postgres"
)

type Storage struct {
	DB *sql.DB
}

func New() (*Storage, error) {
	const op string = "storage.postgres.new"
	cfg := cfg.MustLoad()
	db, err := sql.Open("postgres", cfg.String())
	if err != nil {
		return nil, fmt.Errorf("%s error: %s", op, err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s error: %s", op, err)
	}
	return &Storage{DB: db}, nil
}
