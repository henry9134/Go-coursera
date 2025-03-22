package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Initial slice with length 0 and capacity 3
	nums := make([]int, 0, 3)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter an integer (or 'X' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.EqualFold(input, "X") {
			fmt.Println("Exiting program.")
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer or 'X' to exit.")
			continue
		}

		nums = append(nums, num)
		sort.Ints(nums)
		fmt.Println("Sorted slice:", nums)
	}
}
