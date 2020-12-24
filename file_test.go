package gelpers_test

import (
	"fmt"
	"log"
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

func ExampleReadJSONFile() {
	// The example JSON file contains the following:
	//	{
	//	  "data": {
	//	    "outer_key": {
	//	      "inner_key": {
	//	        "v1": "here's a string",
	//	        "v2": 123
	//	       }
	//	     }
	//	   }
	// 	}

	dataMap, err := gelpers.ReadJSONFile("testdata/example.json")
	if err != nil {
		log.Fatal(err)
	}

	// Now we can do whatever you want with the data map.
	fmt.Println(dataMap)
	// Output: &map[data:map[outer_key:map[inner_key:map[v1:here's a string v2:123]]]]
}

func TestWriteCSVFile(t *testing.T) {
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

func ExampleWriteCSVFile() {
	data := [][]string{
		{"one", "two"},
		{"one_one", "two_two"},
	}

	outfilePath := "testdata/example.csv"
	dataFile, err := gelpers.WriteCSVFile(data, &outfilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(dataFile.Name())

	// Now we can do whatever we want with the output file.
	fileInfo, _ := os.Stat(dataFile.Name())
	fmt.Println(fileInfo.Size())
	// Output: 24
}
