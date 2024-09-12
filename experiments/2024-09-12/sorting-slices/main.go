package main

import (
	"fmt"
	"sort"
)

func uglySort() {
	var digits = []int{8, 6, 7, 5, 3, 0, 9}

	// Create a copy of `digits`
	var sorted = make([]int, len(digits))
	copy(sorted, digits)

	// Now sort the copy
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	fmt.Printf("Digits: %v\n", digits)
	fmt.Printf("Sorted: %v\n", sorted)
}

type LessFunc[T any] func(x, y T) bool

// FYI: This function could go in a shared library
func SortedCopy[T any](src []T, less LessFunc[T]) []T {
	// Create a copy
	var sorted = make([]T, len(src))
	copy(sorted, src)

	// Sort it!
	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

func prettySort() {
	var digits = []int{8, 6, 7, 5, 3, 0, 9}

	var sorted = SortedCopy(digits, func(x, y int) bool {
		return x < y
	})

	fmt.Printf("Digits: %v\n", digits)
	fmt.Printf("Sorted: %v\n", sorted)
}

func main() {
	fmt.Println("uglySort:")
	uglySort()

	fmt.Print("\nprettySort:\n")
	prettySort()
}
