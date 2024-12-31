package tasklist

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func ListAll() {
	db, _ := sql.Open("sqlite", "./sqlite-database.db") // Open the created SQLite File
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal("Task:", err)
	}
	defer rows.Close()
	for rows.Next() {
		fmt.Println(rows)
	}
}
