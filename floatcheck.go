package main

import "fmt"

func HandleFloat() {
	// Now handle the floating-point number input
	var floatInput float64
	fmt.Print("Enter a floating point number: ")
	fmt.Scan(&floatInput)

	// Convert the floating-point number to an integer and print it
	fmt.Println(int(floatInput))
}
