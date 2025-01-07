package dbutils

import (
	"fmt"

	"errors"
	"log"
	"os"

	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func createDb(path string) {
	file, err := os.Create(path) // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
}

func createTable(db *sql.DB) {
	createTasksTableSQL := `CREATE TABLE tasks (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"description" TEXT
		);`

	_, err := db.Exec(createTasksTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Tasks table created")
}

func deleteDb(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}

func validateDb(db *sql.DB) {
	log.Println("Testing the DB")
	// TODO add test to check if the DB is valid when the file exists
}

func testExistence(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		log.Fatal(err)
		return false
	}
}

func InitDb(driver string, path string, reset bool) {
	fmt.Println("init called")

	fmt.Println("Checking if we already have a Db")
	fileExist := testExistence(path)

	if reset && fileExist {
		fmt.Println("Deleting file")
		deleteDb(path)

		fileExist = false
	}

	if fileExist {
		sqliteDatabase, _ := sql.Open(driver, path) // Open the created SQLite File
		defer sqliteDatabase.Close()                // Defer Closing the database
		validateDb(sqliteDatabase)
	} else {
		createDb(path)
		fmt.Println("sqlite-database.db created")

		db := DbConnect(driver, path) // Open the created SQLite File
		defer db.Close()              // Defer Closing the database
		createTable(db)               // Create Database Tables
		//testDb(sqliteDatabase)
	}

}
