/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/atharvwasthere/JustDO/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short: "Mark Items as Done!",
	Long: `Changes the Pending tasks as done and removes from the respective file`,
	Run: doneRun,
}

// doneRun executes the done command functionality
func doneRun(cmd *cobra.Command, args []string) {
	// Read existing items from data file
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		fmt.Println("Could not read the file :(", err)
		return
	}

	// Validate arguments
	if len(args) == 0 {
		log.Fatalln("Please provide a task number")
	}

	// Parse task number
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid task number")
	}

	// Mark task as done if valid index
	if i > 0 && i <= len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Text, "marked done!")

		// Sort and save updated items
		sort.Sort(todo.ByPri(items))
		todo.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "doesn't match any items")
	}
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
