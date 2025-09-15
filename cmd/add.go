/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"log"

	"github.com/atharvwasthere/JustDO/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Priority flag for setting task priority
var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new Todo",
	Long:  `This will add a new todo to the list of Todo's`,
	Run:   addRun,
}

// taskExists checks if a task with the given text already exists in the items slice
func taskExists(items []todo.Item, newText string) bool {
	for _, item := range items {
		if item.Text == newText {
			return true
		}
	}
	return false
}

// addRun is the main function that executes when the add command is called
/*
cmd and args are standard parameters Cobra gives you:
- cmd → current command object
- args → CLI arguments like go run main.go add 1 2
*/
func addRun(cmd *cobra.Command, args []string) {
	var items = []todo.Item{}

	// Read existing items from the data file
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	// Add each new task if it doesn't already exist
	addedCount := 0
	for _, x := range args {
		if !taskExists(items, x) {
			item := todo.Item{Text: x}
			item.SetPriority(priority)
			items = append(items, item)
			addedCount++
		} else {
			fmt.Println("This task already exists :)")
		}
	}

	// Save the updated items list
	err = todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		fmt.Println("Error saving tasks :(")
		fmt.Println(err)
		return
	}

	// Show confirmation message
	if addedCount > 0 {
		totalTasks := len(items)
		pendingTasks := 0
		for _, item := range items {
			if !item.Done {
				pendingTasks++
			}
		}
		fmt.Printf("Added %d task(s). You now have %d task(s) in total (%d pending).\n", addedCount, totalTasks, pendingTasks)
	}
}
func init() {
	rootCmd.AddCommand(addCmd)

	// Here we will define our flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCmd.Flags().IntVarP(&priority, "Priority", "p", 2, "Priority:1,2,3 ")
}
