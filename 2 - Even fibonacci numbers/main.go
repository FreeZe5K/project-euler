package main

import "fmt"

// Entry point for the application
func main() {
	// Create variables to store our total, left, and right numbers
	total, left, right := 0, 1, 2

	// Iterate from 1 to 4 million
	for left < 4e6 {
		// Check if the current value is divisible by two
		if right%2 == 0 {
			// Add the current value to the total
			total += right
		}

		// The next variable is going to be the
		// sum of the left and right variables
		next := left + right

		// The left becomes the right, and the right
		// becomes the next (left + right)
		left, right = right, next
	}

	// Print the total to the console
	fmt.Println(total)
}
