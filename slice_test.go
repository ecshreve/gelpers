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
