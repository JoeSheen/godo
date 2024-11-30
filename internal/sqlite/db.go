package sqlite

import "database/sql"

const driverName = "sqlite3"

type DatabaseContext struct {
	db *sql.DB
}

func OpenDBConnection(fileName string) (*DatabaseContext, error) {
	db, err := sql.Open(driverName, fileName)
	if err != nil {
		return nil, err
	}
	return &DatabaseContext{db: db}, nil
}
