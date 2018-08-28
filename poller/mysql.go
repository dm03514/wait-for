package poller

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(connectionString string) (MySQL, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return MySQL{}, nil
	}

	return MySQL{
		db: db,
	}, nil
}

func (m MySQL) CheckReady() (ready bool, err error) {
	if err = m.db.Ping(); err != nil {
		return false, err
	}
	return true, nil
}
