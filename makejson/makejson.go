package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	nameInput, _ := reader.ReadString('\n')
	nameInput = strings.TrimSpace(nameInput)

	fmt.Print("Enter your address: ")
	addressInput, _ := reader.ReadString('\n')
	addressInput = strings.TrimSpace(addressInput)

	data := map[string]string{
		"name":    nameInput,
		"address": addressInput,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
