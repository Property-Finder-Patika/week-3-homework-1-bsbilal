package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	TRY = iota // iota+1 for begin 1
	EUR
	BTC
)

var (
	rates [3]float64 = [3]float64{17.27, 0.98, 0.000047}
)

func main() {

	fmt.Println(rates[TRY])

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}

	fmt.Printf("%.2f USD is %.2f TRY\n", amount, rates[TRY]*amount)
	fmt.Printf("%.2f TRY is %.2f EUR\n", amount, rates[EUR]*amount)
	fmt.Printf("%.2f BTC is %f BTC\n", amount, rates[BTC]*amount)

}
