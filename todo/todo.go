package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct{
	Text string
}
// Item is exporteddd

func SaveItems(filename string , items []Item) error {
	anytask, err := json.Marshal(items) // returns the JSON encoding of items
	err = os.WriteFile(filename,anytask,0644) //redeclaring so = not  :=
	if err != nil{
		return err
	}
	fmt.Printf(string(anytask))
	return nil
}