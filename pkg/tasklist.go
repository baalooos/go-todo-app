package pkg

import (
	"fmt"
	"log"

	"database/sql"

	"github.com/baalooos/go-todo-app/pkg/dbutils"
	_ "github.com/glebarez/go-sqlite"
)

func ListAll(driver string, path string) {
	var task Task

	db := dbutils.DbConnect(driver, path)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal("Task:", err)
	}
	defer rows.Close()
	for rows.Next() {
		// TODO improve rendering
		// fmt.Println(rows)
		if err := rows.Scan(&task.ID, &task.Name, &task.Description); err != nil {
			if err == sql.ErrNoRows {
				log.Fatal("No task")
			}
			log.Fatal(err)
		}
		fmt.Println(task.ID, task.Name, task.Description)
	}
}

func ListByID(driver string, path string, id int64) {
	var task Task

	db := dbutils.DbConnect(driver, path)
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
