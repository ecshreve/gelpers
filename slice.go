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

// IntMatrixTranspose returns a 2d slice such that input[r][c] = output[c][r].
// It expects a non-nil, non-empty 2d slice where each row has the same length.
func IntMatrixTranspose(m [][]int) [][]int {
	if len(m) == 0 || m == nil {
		return nil
	}

	rowLen := len(m[0])
	for _, row := range m {
		if len(row) != rowLen {
			return m
		}
	}

	var ret [][]int
	for i := 0; i < len(m[0]); i++ {
		var r []int
		for _, entry := range m {
			r = append(r, entry[i])
		}
		ret = append(ret, r)
	}

	return ret
}
