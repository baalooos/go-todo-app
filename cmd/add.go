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

var TaskName string
var TaskDescription string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		var new_task pkg.Task
		new_task.Name = TaskName
		new_task.Description = TaskDescription

		id, err := pkg.TaskAdd(new_task)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Your new task id is:", id)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().StringVarP(&TaskName, "name", "n", "", "Name of the task")
	addCmd.PersistentFlags().StringVarP(&TaskDescription, "description", "d", "", "Description of the task")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
