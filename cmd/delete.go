/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/baalooos/go-todo-app/pkg"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long: `Delete a task using it's ID.
	Example:
	- task delete --id 3
	`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete function called.")
		if Id != "" {
			Ids = append(Ids, Id)
			fmt.Println("delete called with Id:", Id)
			pkg.TaskDelete(Driver, DbPath, Ids)
		} else if len(Ids) > 0 {
			fmt.Println("Deleting multiple tasks")
			pkg.TaskDelete(Driver, DbPath, Ids)

		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().StringVar(&Id, "id", "", "id of the task")
	deleteCmd.PersistentFlags().StringSliceVar(&Ids, "ids", []string{}, "Ids to delete")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
