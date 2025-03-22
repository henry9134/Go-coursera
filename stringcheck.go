package main

import (
	"fmt"
	"strings"
)

func CheckString() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	// Convert input to lowercase to make it case-insensitive
	input = strings.ToLower(input)

	// Check if the string starts with 'i', ends with 'n', and contains 'a'
	if input[0] == 'i' && input[len(input)-1] == 'n' && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
