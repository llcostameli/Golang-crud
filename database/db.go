package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	connection := "root:rootroot@/db_meli"
	db, err := sql.Open("mysql", connection)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
