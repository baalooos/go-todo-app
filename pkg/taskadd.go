package pkg

import (
	"fmt"

	"github.com/baalooos/go-todo-app/pkg/dbutils"
	_ "github.com/glebarez/go-sqlite"
)

func TaskAdd(driver string, path string, task Task) (int64, error) {
	db := dbutils.DbConnect(driver, path)
	defer db.Close()

	fmt.Println(task.Name, task.Description)
	result, err := db.Exec("INSERT INTO tasks (name, description) VALUES (?, ?)", task.Name, task.Description)
	if err != nil {
		return 0, fmt.Errorf("addTask: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
