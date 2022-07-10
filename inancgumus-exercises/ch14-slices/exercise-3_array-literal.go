package main

import (
	"fmt"
	"strings"
)

func main() {

	// friends
	var friends [3]string = [3]string{"Ali", "Veli", "Ayşe"}

	// distances
	distances := [5]float32{.1, .2, 10.0, 20.2, 30.5}

	//data buffer with five bytes
	bytes := [5]byte{'H', 'e', 'l', 'l', 'o'}

	// currency ratios
	ratios := [1]float64{0.02}

	//website alives
	var alives [4]bool = [4]bool{true, false, false, true}

	// empty array
	emptyArray := [0]byte{}
	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("names", separator)
	for i := 0; i < len(friends); i++ {
		fmt.Printf("names[%d]: %q\n", i, friends[i])
	}

	fmt.Print("\ndistances", separator)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Print("\ndata", separator)
	for i := 0; i < len(bytes); i++ {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, bytes[i])
	}

	fmt.Print("\nratios", separator)
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}

	fmt.Print("\nalives", separator)
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	// no loop for zero elements
	fmt.Print("\nzero", separator)
	for i := 0; i < len(emptyArray); i++ {
		fmt.Printf("zero[%d]: %d\n", i, emptyArray[i])
	}

	// =========================================================================

	// you know how this works :) don't be freaked out!
	fmt.Printf(`
%s
FOR RANGES
%[1]s
`, strings.Repeat("~", 30))

	fmt.Print("names", separator)
	for i, v := range friends {
		fmt.Printf("names[%d]: %q\n", i, v)
	}

	fmt.Print("\ndistances", separator)
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

	fmt.Print("\ndata", separator)
	for i, v := range bytes {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, v)
	}

	fmt.Print("\nratios", separator)
	for i, v := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, v)
	}

	fmt.Print("\nalives", separator)
	for i, v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}

	// no loop for zero elements
	fmt.Print("\nzero", separator)
	for i, v := range emptyArray {
		fmt.Printf("zero[%d]: %d\n", i, v)
	}
}
