/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/baalooos/go-todo-app/pkg"
	"github.com/spf13/cobra"
)

var demo_tasks = [3]pkg.Task{
	pkg.Task{
		Name:        "Task1",
		Description: "My first description",
	},
	pkg.Task{
		Name:        "Task2",
		Description: "My second description",
	},
	pkg.Task{
		Name:        "Task3",
		Description: "My third description",
	},
}

// demoCmd represents the demo command
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Add dummy task to your Database",
	Long: `Add dummy task to your Database,
	Example:
	- go-todo-app demo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("demo called")

		for _, task := range demo_tasks {
			pkg.TaskAdd(Driver, DbPath, task)
		}

	},
}

func init() {
	rootCmd.AddCommand(demoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// demoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// demoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
