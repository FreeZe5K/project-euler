package main

import "fmt"

// Entry point for the application
func main() {
	// Create a variable to store a running total
	total := 0

	// Iterate from 0 to 999
	for index := 0; index < 1000; index++ {
		// Check whether the current index is divisible by 3 or 5
		if index%3 == 0 || index%5 == 0 {
			// Add the current index to the total
			total += index
		}
	}

	// Print the total to the console
	fmt.Println(total)
}
