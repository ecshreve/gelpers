package gelpers_test

import (
	"fmt"
	"testing"

	"github.com/ecshreve/gelpers"
	"github.com/stretchr/testify/assert"
)

func TestIntSliceContains(t *testing.T) {
	testcases := []struct {
		desc     string
		inpSlice []int
		inpVal   int
		expected bool
	}{
		{
			desc:     "value exists",
			inpSlice: []int{1, 2, 3},
			inpVal:   2,
			expected: true,
		},
		{
			desc:     "value doesn't exist",
			inpSlice: []int{1, 2, 3},
			inpVal:   999,
			expected: false,
		},
		{
			desc:     "empty slice",
			inpSlice: []int{},
			inpVal:   999,
			expected: false,
		},
		{
			desc:     "nil slice",
			inpSlice: nil,
			inpVal:   999,
			expected: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual := gelpers.IntSliceContains(testcase.inpSlice, testcase.inpVal)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}

func ExampleIntSliceContains() {
	s := []int{1, 2, 3}
	fmt.Println(gelpers.IntSliceContains(s, 2))
	// Output: true
}

func ExampleIntSliceSum() {
	s := []int{1, 2, 3}
	fmt.Println(gelpers.IntSliceSum(s))
	// Output: 6
}

func TestIntSliceMax(t *testing.T) {
	testcases := []struct {
		desc     string
		inp      []int
		expected int
	}{
		{
			desc:     "simple",
			inp:      []int{1, 2, 3},
			expected: 3,
		},
		{
			desc:     "single item slice",
			inp:      []int{1},
			expected: 1,
		},
		{
			desc:     "empty slice",
			inp:      []int{},
			expected: -1,
		},
		{
			desc:     "nil slice",
			inp:      nil,
			expected: -1,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual := gelpers.IntSliceMax(testcase.inp)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}

func TestIntSliceMin(t *testing.T) {
	testcases := []struct {
		desc     string
		inp      []int
		expected int
	}{
		{
			desc:     "simple",
			inp:      []int{1, 2, 3},
			expected: 1,
		},
		{
			desc:     "single item slice",
			inp:      []int{1},
			expected: 1,
		},
		{
			desc:     "empty slice",
			inp:      []int{},
			expected: -1,
		},
		{
			desc:     "nil slice",
			inp:      nil,
			expected: -1,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual := gelpers.IntSliceMin(testcase.inp)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}

func ExampleIntSliceMax() {
	s := []int{1, 2, 3}
	fmt.Println(gelpers.IntSliceMax(s))
	// Output: 3
}

func ExampleIntSliceMin() {
	s := []int{1, 2, 3}
	fmt.Println(gelpers.IntSliceMin(s))
	// Output: 1
}
