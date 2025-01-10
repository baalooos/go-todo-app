package pkg

import (
	"fmt"

	"github.com/baalooos/go-todo-app/pkg/dbutils"
	_ "github.com/glebarez/go-sqlite"
)

func TaskAdd(driver string, path string, task Task) error {
	var taskList []Task

	db := dbutils.DbConnect(driver, path)
	defer db.Close()

	fmt.Println(task.Name, task.Description)
	result, err := db.Exec("INSERT INTO tasks (name, description) VALUES (?, ?)", task.Name, task.Description)
	if err != nil {
		return fmt.Errorf("addTask: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("addAlbum: %v", err)
	}
	task.ID = id

	taskList = append(taskList, task)
	fmt.Println("New task added:")
	PrintTasks(taskList)

	return nil
}
