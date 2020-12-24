package gelpers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

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
