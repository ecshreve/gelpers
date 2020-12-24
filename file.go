package gelpers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/samsarahq/go/oops"
)

// ReadJSONFile returns a pointer to a map representation of the JSON file at
// the given path, or an error if unsuccessful.
func ReadJSONFile(path string) (*map[string]interface{}, error) {
	// Check that the given file is a JSON file.
	ext := filepath.Ext(path)
	if ext != ".json" {
		return nil, oops.Errorf("file at path: %s is not a JSON file", path)
	}

	// Open the file specified by path.
	file, err := os.Open(path)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to open file at path: %s", path)
	}
	defer file.Close()

	// Read the file into a byte array.
	//
	// The error case here is not covered by unit tests because it would involve
	// mocking the filesystem and doing some weird things to force an error.
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to read file: %s to byte array", file.Name())
	}

	// Unmarshall the byte array into a map.
	var result map[string]interface{}
	if err = json.Unmarshal([]byte(byteValue), &result); err != nil {
		return nil, oops.Wrapf(err, "unable to unmarshal byte array to map for file: %s", file.Name())
	}

	return &result, nil
}

// WriteCSVFile writes the given 2d slice of strings to a CSV file at the given
// path and returns a pointer to the output file, or an error if unsuccessful.
//
// This function treats the first row in the data argument as the headers for
// the CSV file.
//
// If no path is provided, then a default filename is generated for the output.
// The default directory for the output file is the root directory of the
// module. The filename is of the form `data_<ms_epoch>.output.csv`.
//
// If a path is provided it must be for a CSV file.
func WriteCSVFile(data [][]string, path *string) (*os.File, error) {
	var outfilePath *string
	if path == nil || *path == "" {
		timeNowMs := time.Now().Unix()
		outfilePathVal := fmt.Sprintf("data_%d.output.csv", timeNowMs)
		outfilePath = &outfilePathVal
	} else {
		ext := filepath.Ext(*path)
		if ext != ".csv" {
			return nil, oops.Errorf("output file must be a CSV file: %s", *path)
		}
		outfilePath = path
	}

	// Create the output file.
	file, err := os.Create(*outfilePath)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to create output file for path: %s", *outfilePath)
	}
	defer file.Close()

	// Create a CSV writer, and write the data to the output file.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to write value: %+v to file: %s", value, file.Name())
		}
	}

	return file, nil
}
