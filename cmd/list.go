/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/atharvwasthere/JustDO/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Command line flags for list filtering
var (
	doneOpt bool
	allOpt  bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the Todo's",
	Long:  `When run performs listing of all the Todo's`,
	Run:   listRun,
}

// listRun executes the list command functionality
func listRun(cmd *cobra.Command, args []string) {
	// Read items from data file
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
		return
	}

	// Sort items by priority and update positions
	sort.Sort(todo.ByPri(items))
	for i := range items {
		items[i].Position = i + 1
	}

	// Create tabwriter for formatted output
	// func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	// Display items based on filter flags
	for _, item := range items {
		if allOpt || item.Done == doneOpt {
			fmt.Fprintln(w, item.Label()+"\t"+item.PrettyDone()+"\t"+item.PrettyP()+"\t"+item.Text+"\t")
		}
	}

	// Flush ensures any incomplete or buffered content is treated as whole and formatted
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt,"done",false,"Show 'Done' Todos")
	listCmd.Flags().BoolVar(&allOpt,"all",false,"Show 'All' Todos" )

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
