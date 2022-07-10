package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// the program awaits input from user. that works with go pxxx.go "number"
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	// the number parses from input values
	number, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid input. It should be a number.")
		return
	}

	// results comes from functions
	fmt.Printf("Prime factors of  %d %v", number, PrimeFactors(number))
}

func PrimeFactors(n int) []int {

	// that function returns  prime factors of given number
	var (
		factors []int
	)
	divisor := 2
	for n >= 2 {
		//divisor begins from 2 and recursed itself
		if n%divisor == 0 {
			factors = append(factors, divisor)
			n = n / divisor
		} else {
			divisor++
		}
	}
	return factors
}
