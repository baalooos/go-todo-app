/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/baalooos/go-todo-app/pkg"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your tasks",
	Long: `List your tasks.
	It is possible to list either all your tasks or only one task.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		if All {
			pkg.ListAll(Driver, DbPath)
		} else if Id != "" {
			Ids = append(Ids, Id)
			pkg.ListByID(Driver, DbPath, Ids)
		} else if len(Ids) > 0 {
			pkg.ListByID(Driver, DbPath, Ids)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().BoolVarP(&All, "all", "a", false, "List all tasks")
	listCmd.PersistentFlags().StringVar(&Id, "id", "", "id of the task")
	listCmd.PersistentFlags().StringSliceVar(&Ids, "ids", []string{}, "Ids to delete")
	listCmd.MarkFlagsOneRequired("all", "id", "ids")
}
