package gelpers_test

import (
	"fmt"
	"testing"

	"github.com/ecshreve/gelpers"
	"github.com/stretchr/testify/assert"
)

func TestSnakeToCamel(t *testing.T) {
	testcases := []struct {
		desc     string
		inp      string
		expected string
	}{
		{
			desc:     "simple",
			inp:      "foo_bar_baz",
			expected: "FooBarBaz",
		},
		{
			desc:     "no underscores",
			inp:      "foobarbaz",
			expected: "Foobarbaz",
		},
		{
			desc:     "input contains numbers",
			inp:      "foo_bar2_baz3",
			expected: "FooBar2Baz3",
		},
		{
			desc:     "empty",
			inp:      "",
			expected: "",
		},
		{
			desc:     "input contains spaces",
			inp:      "foo bar baz",
			expected: "foo bar baz",
		},
		{
			desc:     "input starts with a number",
			inp:      "999foo_bar_baz",
			expected: "999foo_bar_baz",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual := gelpers.SnakeToCamel(testcase.inp)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}

func ExampleSnakeToCamel() {
	fmt.Println(gelpers.SnakeToCamel("foo_bar_baz"))
	// Output: FooBarBaz
}

func TestSpaceSepToCamel(t *testing.T) {
	testcases := []struct {
		desc     string
		inp      string
		expected string
	}{
		{
			desc:     "simple",
			inp:      "foo bar baz",
			expected: "FooBarBaz",
		},
		{
			desc:     "no spaces",
			inp:      "foobarbaz",
			expected: "Foobarbaz",
		},
		{
			desc:     "input contains numbers",
			inp:      "foo bar2 baz3",
			expected: "FooBar2Baz3",
		},
		{
			desc:     "empty",
			inp:      "",
			expected: "",
		},
		{
			desc:     "input contains underscores",
			inp:      "foo_bar baz",
			expected: "foo_bar baz",
		},
		{
			desc:     "input starts with a number",
			inp:      "999foo bar baz",
			expected: "999foo bar baz",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual := gelpers.SpaceSepToCamel(testcase.inp)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}

func ExampleSpaceSepToCamel() {
	fmt.Println(gelpers.SpaceSepToCamel("foo bar baz"))
	// Output: FooBarBaz
}
