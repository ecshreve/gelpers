package gelpers

import "sort"

// IntSliceContains checks if the given slice contains the given value.
func IntSliceContains(s []int, v int) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}
	return false
}

// IntSliceSum returns the sum of all items in the given slice of ints.
func IntSliceSum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

// IntSliceMax returns the largest value in the given slice of ints, or -1 if
// the slice is empty or nil.
func IntSliceMax(s []int) int {
	if len(s) < 1 || s == nil {
		return -1
	}

	sort.Ints(s)
	return s[len(s)-1]
}

// IntSliceMin returns the smallest value in the given slice of ints, or -1 if
// the slice is empty or nil.
func IntSliceMin(s []int) int {
	if len(s) < 1 || s == nil {
		return -1
	}

	sort.Ints(s)
	return s[0]
}
