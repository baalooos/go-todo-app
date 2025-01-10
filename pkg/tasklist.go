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
	var taskList []Task

	db := dbutils.DbConnect(driver, path)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal("Task:", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&task.ID, &task.Name, &task.Description); err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No task")
			} else {
				log.Fatal(err)
			}
		} else {
			taskList = append(taskList, task)
		}
	}
	PrintTasks(taskList)
}

func ListByID(driver string, path string, ids []string) {
	var task Task
	var taskList []Task

	db := dbutils.DbConnect(driver, path)
	defer db.Close()

	for _, id := range ids {
		// TODO make it only one call
		row := db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
		if err := row.Scan(&task.ID, &task.Name, &task.Description); err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No task")
			} else {
				log.Fatal(err)
			}
		} else {
			taskList = append(taskList, task)
		}
	}
	PrintTasks(taskList)
}
