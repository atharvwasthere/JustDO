package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct{
	Text string
	Priority int
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
	return items , nil
	
}