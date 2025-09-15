package todo

import (
	"encoding/json"
	"os"
	"strconv"
)

// ByPri implements sort.Interface for []Item based on Priority field
type ByPri []Item

// Item represents a single todo item with text, priority, position and completion status
type Item struct {
	Text     string
	Priority int
	Position int
	Done     bool
}

// Sort interface implementation for ByPri
func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	// Done items come first (when sorted)
	if s[i].Done != s[j].Done {
		return s[i].Done
	}
	// If same done status, sort by priority
	if s[i].Priority == s[j].Priority {
		return s[i].Position < s[j].Position
	}
	return s[i].Priority < s[j].Priority
}

// SaveItems writes the items slice to a JSON file
func SaveItems(filename string, items []Item) error {
	// Marshal items to JSON
	anytask, err := json.Marshal(items) // returns the JSON encoding of items
	if err != nil {
		return err // marshal error
	}

	// Write JSON to file
	err = os.WriteFile(filename, anytask, 0644) // redeclaring so = not :=
	if err != nil {
		return err // writefile err
	}

	return nil
}

// SetPriority sets the priority level for an item
// Priority levels: 1 (low), 2 (medium/default), 3 (high)
func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1: // Low priority - can procrastinate
		i.Priority = 1
	case 3: // High priority
		i.Priority = 3
	default: // Medium priority if not specified
		i.Priority = 2
	}
}

// PrettyP returns a formatted priority indicator
func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 3 {
		return "(3)"
	}
	return " "
}

// PrettyDone returns a visual indicator for completed tasks
func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	} else {
		return ""
	}
}

// Label returns the formatted position label for display
func (i *Item) Label() string {
	return strconv.Itoa(i.Position) + "."
}

// ReadItems reads and parses items from a JSON file
func ReadItems(filename string) ([]Item, error) {
	// Read file contents
	anytask, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}

	// Unmarshal JSON into items slice
	var items []Item
	err = json.Unmarshal(anytask, &items)
	if err != nil {
		return []Item{}, err
	}

	// Update positions for display
	for i := range items {
		items[i].Position = i + 1 // i is basically index here
	}

	return items, nil
}
