package db_init

import (
	"fmt"

	//"database/sql"
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
		"description" TEXT,
		"creation_date" DATE,
		"due_date" DATE
		);` // SQL Statement for Create Table

	log.Println("Create tasks table...")
	statement, err := db.Prepare(createTasksTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")
}

func testDb(db *sql.DB) {
	testDb := `SELECT * FROM tasks;`
	log.Println("Testing the DB")
	statement, err := db.Prepare(testDb)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()

}

func DbInit() {
	if _, err := os.Stat("sqlite-database.db"); err == nil {
		fmt.Println("File already exists")
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
		testDb(sqliteDatabase)
	}

}
