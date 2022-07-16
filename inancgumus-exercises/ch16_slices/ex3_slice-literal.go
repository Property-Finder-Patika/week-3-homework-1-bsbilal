package main

import "fmt"

func main() {
	var (
		names     []string  // The names of your friends
		distances []int     // The distances
		data      []byte    // A data buffer
		ratios    []float64 // Currency exchange ratios
		alives    []bool    // Up/Down status of web servers
	)

	names = []string{"Ali", "Veli", "Ayşe"}
	distances = []int{1, 2, 3}
	data = []byte{'a', 'b', 'c'}
	ratios = []float64{1.4, .2, 8.7}
	alives = []bool{true, false, true, false}

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)

	if len(distances) == len(data) {
		fmt.Println("The length of the distances  the data slices are the same. That means they has same element count.")
	}

}
