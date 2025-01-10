/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/baalooos/go-todo-app/pkg"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Long: `Add a task. You can either do it using flags or dynamically in your terminal.
	Examples:
	- task add
	- task add -n "task_name" -d "my task description"
	`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		var new_task pkg.Task

		new_task.Name = TaskName
		new_task.Description = TaskDescription

		err := pkg.TaskAdd(Driver, DbPath, new_task)
		if err != nil {
			log.Fatal(err.Error())
		}

		// TODO add interactive command line
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().StringVarP(&TaskName, "name", "n", "", "Name of the task")
	addCmd.PersistentFlags().StringVarP(&TaskDescription, "description", "d", "", "Description of the task")
	addCmd.MarkFlagsRequiredTogether("name", "description")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
