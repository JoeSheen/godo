package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const driverName = "sqlite3"

func NewSqlDB(fileName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, fileName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
