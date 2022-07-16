package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		averageArr [5]float64
		sum        float64
	)
	input := os.Args[1:]

	// i do not know but its begins with [1] fmt.Println(len(os.Args))
	if len(input) != 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers)..")
		return
	}

	for i := range input {
		arg, err := strconv.ParseFloat(input[i], 64)
		if err == nil {
			averageArr[i] = arg
			sum += arg
		} else {
			averageArr[i] = 0.0
		}

	}

	fmt.Println("Your entered numbers were", input)
	fmt.Printf("\nYour entered numbers average is %.2f", sum/5)

}
