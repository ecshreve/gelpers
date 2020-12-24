package gelpers_test

import (
	"os"
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

func TestWriteCSVFile(t *testing.T) {
	// This is a valid 2d slice of strings.
	data := [][]string{
		{"one", "two"},
		{"one_one", "two_two"},
	}

	testcases := []struct {
		description string
		outfilePath string
		expectError bool
	}{
		{
			description: "expect error for a bad file type",
			outfilePath: "testdata/nonexistentdirectory/testoutput.xls",
			expectError: true,
		},
		{
			description: "expect error for a bad file path",
			outfilePath: "testdata/nonexistentdirectory/testoutput.csv",
			expectError: true,
		},
		{
			description: "expect success for a good path",
			outfilePath: "testdata/testoutput.csv",
			expectError: false,
		},
		{
			description: "expect success with no path provided",
			outfilePath: "",
			expectError: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			outfile, err := gelpers.WriteCSVFile(data, &testcase.outfilePath)
			assert.Equal(t, testcase.expectError, err != nil)
			assert.Equal(t, testcase.expectError, outfile == nil)

			// Remove the CSV file if one was created.
			if !testcase.expectError {
				os.Remove(outfile.Name())
			}
		})
	}
}
