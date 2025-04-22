package main

import (
	"fmt"
	"sort"
)

// Define a type constraint for ordered types
type Ordered interface {
	~int | ~int64 | ~float64
}

// Generic median function
func median[T Ordered](data []T) float64 {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	n := len(data)
	if n%2 == 0 {
		mid1 := float64(data[n/2-1])
		mid2 := float64(data[n/2])
		return (mid1 + mid2) / 2
	} else {
		return float64(data[n/2])
	}
}

func main() {
	odd := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	even := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("median for odd:", median(odd))   // 5
	fmt.Println("median for even:", median(even)) // 5.5
}

// go install golang.org/x/exp/constraints@latest
// go run median.go
