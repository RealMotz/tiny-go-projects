package main

import (
	"fmt"
	"os"
)

func main() {
	Bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(Bookworms)
}
