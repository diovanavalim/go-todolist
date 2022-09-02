package database

import (
	"database/sql"
	"todolist/cmd/api/config"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionString)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
