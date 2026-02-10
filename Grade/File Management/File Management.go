package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Subject structure to store course details
type Subject struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
	Grade string  `json:"grade"`
}

// SaveGrades writes the data to a JSON file
func SaveGrades(filename string, data []Subject) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, file, 0644)
}

func main() {
	fmt.Println("--- Grade Management System ---")
	// Your logic here...
}
