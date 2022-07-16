package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Fix
//
//  1. Uncomment the code
//
//  2. And fix the problems
//
//  3. BONUS: Simplify the code
// ---------------------------------------------------------

func main() {
	var names [3]string = [3]string{
		"Einstein",
		"Shepard",
		"Tesla"}

	var books [5]string = [5]string{
		"Kafka's Revenge",
		"Stay Golden",
		"",
		"",
		""}
	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
	// simplified

	fmt.Println("Simplified section")
	namess := [3]string{"Einstein", "Shepard", "Tesla"}
	bookss := [5]string{"Kafka's Revenge", "Stay Golden", "", "", ""}
	fmt.Printf(" %#v\n", namess)
	fmt.Printf(" %#v\n", bookss)

}
