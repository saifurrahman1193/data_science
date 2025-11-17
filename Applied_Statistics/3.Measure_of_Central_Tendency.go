package main

import (
	"fmt"
	"sort"
)

func mean(nums []float64) float64 {
	total := 0.0
	for _, n := range nums {
		total += n
	}
	return total / float64(len(nums))
}

func median(nums []float64) float64 {
	sort.Float64s(nums)
	mid := len(nums) / 2
	if len(nums)%2 == 0 {
		return (nums[mid-1] + nums[mid]) / 2
	}
	return nums[mid]
}

func mode(nums []float64) []float64 {
	freq := make(map[float64]int)
	maxFreq := 0
	for _, n := range nums {
		freq[n]++
		if freq[n] > maxFreq {
			maxFreq = freq[n]
		}
	}

	var modes []float64
	for n, f := range freq {
		if f == maxFreq {
			modes = append(modes, n)
		}
	}

	// If every number appears once, there's no mode
	if len(modes) == len(freq) {
		return []float64{}
	}
	return modes
}

func main() {
	data := []float64{1, 2, 2, 3, 4, 4, 4, 5, 6}

	fmt.Printf("Data: %v\n", data)
	fmt.Printf("Mean: %.2f\n", mean(data))
	fmt.Printf("Median: %.2f\n", median(data))
	fmt.Printf("Mode: %v\n", mode(data))
}
