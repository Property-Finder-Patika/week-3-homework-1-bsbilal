package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		sortArr [5]float64
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
			sortArr[i] = arg
		} else {
			sortArr[i] = 0.0
		}

	}
	fmt.Println("Your entered numbers were", sortArr)

	NumberSort(&sortArr)

	fmt.Printf("\nYour sorted numbers are  %.2f", sortArr)

}

func NumberSort(array *[5]float64) {

	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}
