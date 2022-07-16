package main

import "fmt"

func main() {
	var (
		names     []int
		distances []int
		data      []byte
		ratios    []float32
		alives    []bool
	)

	fmt.Printf("names: %T len : %d  is nill %t \n", names, len(names), names == nil)
	fmt.Printf("distances: %T len : %d  is nill %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data: %T len : %d is nill %t \n", data, len(data), data == nil)
	fmt.Printf("ratios: %T len : %d is nill %t \n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives: %T len : %d is nill %t \n", alives, len(alives), alives == nil)

}
