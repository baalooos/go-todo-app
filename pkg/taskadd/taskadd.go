package taskadd

import (
	"fmt"

	"database/sql"

	"github.com/baalooos/go-todo-app/pkg/object"
	_ "github.com/glebarez/go-sqlite"
)

func TaskAdd(task object.Task) (int64, error) {
	db, _ := sql.Open("sqlite", "./sqlite-database.db") // Open the created SQLite File
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
