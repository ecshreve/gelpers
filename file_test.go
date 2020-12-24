package gelpers_test

import (
	"testing"

	"github.com/ecshreve/gelpers"
	"github.com/stretchr/testify/assert"
)

func TestReadJSONFile(t *testing.T) {
	testcases := []struct {
		description string
		infilePath  string
		expectError bool
	}{
		{
			description: "expect error if infile isn't a json file",
			infilePath:  "testdata/filename.somenonjsonextension",
			expectError: true,
		},
		{
			description: "expect error if infile doesn't exist",
			infilePath:  "testdata/nonexistentfile.json",
			expectError: true,
		},
		{
			description: "expect error if infile is malformed",
			infilePath:  "testdata/jsontest_bad.json",
			expectError: true,
		},
		{
			description: "expect success for valid infile",
			infilePath:  "testdata/jsontest.json",
			expectError: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			rawData, err := gelpers.ReadJSONFile(testcase.infilePath)
			if testcase.expectError {
				assert.Error(t, err)
				assert.Nil(t, rawData)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, rawData)
			}
		})
	}
}
