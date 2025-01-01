package pkg

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func ListAll() {
	db := DbConnect()
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

func ListByID(id int64) {
	var task Task

	db := DbConnect()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
	if err := row.Scan(&task.ID, &task.Name, &task.Description); err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No task")
		}
		log.Fatal(err)
	}
	fmt.Println(task.ID, task.Name, task.Description)
}
