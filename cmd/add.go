/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"
	"fmt"

	"github.com/atharvwasthere/LearningGO/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new Todo",
	Long: `This will add a new todo to the list of Todo's`,
	Run: addRun,
	// Run: func(cmd *cobra.Command, args []string) {

	// 	for _,x := range args{
	// 		fmt.Println(x)
	// 	}
	// },
}
/* 
cmd and args are standard parameters Cobra gives you:

    cmd → current command object

    args → CLI arguments like go run main.go add 1 2
*/
// new functions are always wriiten outside var addCmd 
func addRun(cmd *cobra.Command, args []string){
	var items = []todo.Item{} 
	for _,x := range args{
		items = append(items , todo.Item{Text: x}) 
	}
	// fmt.Println(items)
	// fmt.Printf("%#v\n",items)
	err := todo.SaveItems("C:/Users/athar/desktop/CS/Golang/Learning CLI/tasks.json", items);
	if err != nil{
		fmt.Errorf("%v",err)
	}
	// todo.SaveItems("x",items)
}
func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
