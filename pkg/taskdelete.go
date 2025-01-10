package pkg

import (
	"fmt"

	"github.com/baalooos/go-todo-app/pkg/dbutils"
	_ "github.com/glebarez/go-sqlite"
)

func TaskDelete(driver string, path string, ids []string) {
	db := dbutils.DbConnect(driver, path)
	defer db.Close()

	for _, id := range ids {
		res, err := db.Exec("DELETE FROM tasks WHERE id=$1", id)

		if err == nil {
			count, err := res.RowsAffected()
			if err == nil {
				if count == 1 {
					fmt.Println("One task deleted.")
				} else if count == 0 {
					fmt.Println("No task deleted")
				}
			}

		} else {
			fmt.Println(err)
		}
	}
}
