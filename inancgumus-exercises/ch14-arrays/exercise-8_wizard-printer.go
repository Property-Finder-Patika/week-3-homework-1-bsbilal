package main

import (
	"fmt"
	"strings"
)

func main() {
	wizards := [6][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "emc2"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"}}

	for i := range wizards {

		fmt.Printf("%-15s %-15s %-15s\n", wizards[i][0], wizards[i][1], wizards[i][2])
		if i == 0 {
			fmt.Println(strings.Repeat("=", 50))
		}
	}

}
