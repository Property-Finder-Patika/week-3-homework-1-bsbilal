package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str := "slice rotates"
	sliceRotates := strings.Split(str, "")

	fromLeft := false

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Please provide 2 keys as like:'r 2' that shows result.\nr means right and l means left and the second is begin space for string.")
		return
	}

	// side defines the rotate
	side := args[0]
	if side != "l" {
		fromLeft = true
	} else if side != "r" {
		fromLeft = false

	} else {
		fmt.Println("Invalid input. It should be a r or l for first element.")
		return
	}

	// the number means begin index of rotate
	number, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid input. It should be a number.")
		return
	}

	fmt.Println("Main string was : ", sliceRotates)
	fmt.Printf("That rotate begins from : %d \n", number)
	if fromLeft {
		fmt.Println("It begins from left side.")
	} else {
		fmt.Println("It begins from right side.")
	}

	WriteSlicedRotates(fromLeft, number, sliceRotates)

}
func WriteSlicedRotates(isLeft bool, begin int, word []string) {
	// for structure that function created
	beginIndex := 0
	if isLeft {
		beginIndex = begin
	} else {
		beginIndex = len(word) - 1 - begin

		//reverse function
		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			word[i], word[j] = word[j], word[i]
		}
	}

	fmt.Printf("Then new string look like : %v ", word[beginIndex:])

}
