package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// A Bookworm contains the list of books on a bookworm's shelf.
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// loadBookworms reads the file and returns the list of bookworms, and their beloved books, found therein.
func loadBookworms(filePath string) ([]Bookworm, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	var bookworms []Bookworm
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&bookworms)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v", err)
		return nil, err
	}

	return bookworms, nil
}
