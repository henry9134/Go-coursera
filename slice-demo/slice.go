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
	var nums []int
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter an integer (or 'X' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.EqualFold(input, "X") {
			fmt.Println("Exiting.")
			break
		}

		if num, err := strconv.Atoi(input); err == nil {
			nums = append(nums, num)
			sort.Ints(nums)
			fmt.Println("Sorted:", nums)
		} else {
			fmt.Println("Invalid input.")
		}
	}
}
