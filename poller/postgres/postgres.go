package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func New(connectionString string) (Postgres, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return Postgres{}, err
	}
	return Postgres{
		db: db,
	}, nil
}

func (p Postgres) CheckReady() (ready bool, err error) {
	if err = p.db.Ping(); err != nil {
		return false, err
	}
	return true, nil
}
