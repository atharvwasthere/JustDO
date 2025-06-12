/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"
	"fmt"
	"log"

	"github.com/atharvwasthere/JustDO/todo"
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

var priority int
/* 
cmd and args are standard parameters Cobra gives you:

    cmd → current command object

    args → CLI arguments like go run main.go add 1 2
*/
// new functions are always wriiten outside var addCmd 

func taskExists(items []todo.Item, newText string ) bool{
	for _,item :=  range items{
		if item.Text == newText{
			return true
		}
	}
	return false
}


func addRun(cmd *cobra.Command, args []string){
	var items = []todo.Item{} 
	//"C:/Users/athar/desktop/CS/Golang/Learning CLI/tasks.json"
	items , err := todo.ReadItems(datafile)

	if err != nil{
		log.Printf("%v",err)
	}
	for _,x := range args{
		if !taskExists(items,x){
			item := todo.Item{Text:x}
			item.SetPriority(priority)
			items = append(items , item) 
		}else{
			fmt.Println("This task already exists :) ")
		}
	}

	// fmt.Println(items)
	// fmt.Printf("%#v\n",items)
	err = todo.SaveItems(datafile, items);
	if err != nil{
		fmt.Println("Error saving tasks :(")
		fmt.Println(err)
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

	addCmd.Flags().IntVarP(&priority, "Priority", "p" , 2  ,"Priority:1,2,3 " )
}
