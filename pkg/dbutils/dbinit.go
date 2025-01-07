package pkg

import (
	"fmt"

	"errors"
	"log"
	"os"

	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func createTable(db *sql.DB) {
	createTasksTableSQL := `CREATE TABLE tasks (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"description" TEXT
		);`
	// "creation_date" DATE,
	// "due_date" DATE
	// SQL Statement for Create Table

	log.Println("Create tasks table...")
	_, err := db.Exec(createTasksTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Tasks table created")
}

func testDb(db *sql.DB) {
	log.Println("Testing the DB")

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal("Task:", err)
	}
	defer rows.Close()
	for rows.Next() {
		fmt.Println(rows)
	}

}

func InitDb() {
	if _, err := os.Stat("sqlite-database.db"); err == nil {
		fmt.Println("File already exists")

		sqliteDatabase, _ := sql.Open("sqlite", "./sqlite-database.db") // Open the created SQLite File
		defer sqliteDatabase.Close()                                    // Defer Closing the database
		testDb(sqliteDatabase)
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("init called")
		file, err := os.Create("sqlite-database.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		fmt.Println("sqlite-database.db created")

		sqliteDatabase, _ := sql.Open("sqlite", "./sqlite-database.db") // Open the created SQLite File
		defer sqliteDatabase.Close()                                    // Defer Closing the database
		createTable(sqliteDatabase)                                     // Create Database Tables
		//testDb(sqliteDatabase)
	}

}
