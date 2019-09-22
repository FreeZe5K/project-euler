package main

import "fmt"

// Entry point for the application
func main() {
	// Create variables to store our current
	// factor and the value we want to test
	factor, val := 3, 600851475143

	// Iterate from the factor to the value
	for factor < val {
		// Check whether the current factor goes
		// into the value evenly
		if val%factor == 0 {
			// The value will be set to the current
			// value divided by the current factor
			val /= factor
		} else {
			// Increase the factor by two steps
			factor += 2
		}
	}

	// Print the total to the console
	fmt.Println(factor)
}
