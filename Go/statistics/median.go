package main

import "fmt"

func median(data []int) int {
	total := 0
	for _, item := range data {
		total += item
	}
	return total / len(data)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("median:", median(numbers))
}

// go run median.go
