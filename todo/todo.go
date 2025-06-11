package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Item struct{
	Text string
	Priority int
	Position int
}
// Item is exporteddd

func SaveItems(filename string , items []Item) error {
	anytask, err := json.Marshal(items) // returns the JSON encoding of items
	if err !=nil{
		return err  // marshal error
	}
	err = os.WriteFile(filename,anytask,0644) //redeclaring so = not  :=
	if err != nil{
		return err // writefile err
	}
	fmt.Printf(string(anytask))
	return nil
}

func  (i*Item) SetPriority(pri int){
	switch pri {
	// can procastinate	
	case 1 :
		i.Priority = 1
	// high
	case 3 :
		i.Priority = 3
	// med if not specified 
	default :
		i.Priority = 2

	}
}
func  (i*Item) PrettyP() string{
	if i.Priority == 1{
		return "(1)"
	}
	if i.Priority == 3{
		return "(3)"
	}
	return " "
}

func (i*Item) Label() string {
	return strconv.Itoa(i.Position) + "."
}



func ReadItems(filename string )  ([]Item , error) {
	anytask, err := os.ReadFile(filename)
	if err != nil{
		return []Item {} , err
	}
	var items []Item
	err = json.Unmarshal(anytask ,&items);
	if err != nil{
		return []Item {} , err
	}
	for i,_ := range items{
		items[i].Position = i + 1 //  i is basically index here 
	}

	return items , nil
	
}