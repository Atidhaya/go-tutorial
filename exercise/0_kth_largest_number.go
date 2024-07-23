package exercise

import (
	"fmt"
	"sort"
)

// Instruction
// implement and write test both happy and bad cases.

func KthLargestNumber(numbers []int, k int) (int, error) {
	if k <= 0 {
		return 0, fmt.Errorf("k must be positive integer")
	}
	if k > len(numbers) {
		return 0, fmt.Errorf("k cannot exceed array length")
	}
	// Write your code here.
	// https://pkg.go.dev/sort
	// this is why I told you golang has bad DX
	sorted := sort.IntSlice(numbers)
	return sorted[len(sorted)-k], nil
}
