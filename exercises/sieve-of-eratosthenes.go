package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please provide just integer number.")
		return
	}

	// the number comes from go run side
	number, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid input. It should be a number.")
		return
	}

	// the number slots array stores empty boolean values which has returns true after prime value checks
	var numberSlots []bool
	numberSlots = make([]bool, number)

	for i := number - 1; i > 1; i-- {
		isDivided := false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isDivided = true
				break
			}
		}
		if !isDivided {
			// that number did not divide any number so that that is prime
			numberSlots[i] = true
		}
	}

	// we gemerated a mew number for get indices of prime numbers
	var primeNumbers []int
	for i := 0; i < number; i++ {
		if numberSlots[i] {
			primeNumbers = append(primeNumbers, i)
		}

	}

	// we will shows the prime numbers
	fmt.Println(primeNumbers)
}
