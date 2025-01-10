package pkg

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintTasks(taskList []Task) {
	// Create a new tabwriter.Writer instance.
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)

	fmt.Fprintln(w, "Id\tName\tDescription")
	for _, task := range taskList {
		fmt.Fprintln(w, task.ID, "\t", task.Name, "\t", task.Description)
	}

	w.Flush()
}
