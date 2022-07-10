package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	books = [...]string{"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition"}
)

func main() {

	input := os.Args[1:]
	if len(input) != 1 {
		fmt.Println("Tell me a book title.")
		return
	}

	found := false
	fmt.Printf("Search result of %s: \n", input[0])
	for _, book := range books {
		if strings.Contains(strings.ToLower(book), input[0]) {
			found = true
			fmt.Printf("+ %s \n", book)
		}
	}
	if !found {
		fmt.Printf("\n We do not have book: %s", input[0])

	}

}
