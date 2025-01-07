package dbutils

import (
	"log"

	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func DbConnect(driver string, path string) *sql.DB {
	db, err := sql.Open(driver, path)
	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}
	return db
}
