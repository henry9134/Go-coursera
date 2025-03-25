package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	fmt.Print("Enter the name of the file: ")
	var filename string
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var names []name
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) >= 2 {
			entry := name{
				fname: parts[0],
				lname: parts[1],
			}
			names = append(names, entry)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, n := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", n.fname, n.lname)
	}
}
