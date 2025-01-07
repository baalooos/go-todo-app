package pkg

import (
	"log"

	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func DbConnect() *sql.DB {
	db, err := sql.Open("sqlite", "./sqlite-database.db")
	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}
	return db
}
