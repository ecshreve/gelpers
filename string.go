package gelpers

import (
	"bytes"
	"log"
	"regexp"
	"strings"
)

// SnakeToCamel converts a snake_case_string to a CamelCaseString. If the input
// is not valid then it returns the input string.
//
// A string is considered valid if it satisfies the following:
// 	- non-empty
// 	- begins with a letter
//	- only contains letters, numbers, and the underscores
func SnakeToCamel(s string) string {
	validSnakeRe := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]*(_[a-zA-Z0-9]+)*$`)
	if !validSnakeRe.MatchString(s) {
		log.Println("input is not a valid snake case string")
		return s
	}

	b := []byte(strings.TrimSpace(s))
	buffer := make([]byte, 0, len(s))

	tokens := bytes.Split(b, []byte("_"))
	for _, tok := range tokens {
		buffer = append(buffer, bytes.ToUpper(tok[:1])...)
		buffer = append(buffer, tok[1:]...)
	}

	return string(buffer)
}

// SpaceSepToCamel converts a space separated string to a CamelCaseString. If
// the input is not valid then it returns the input string.
//
// A string is considered valid if it satisfies the following:
// 	- non-empty
// 	- begins with a letter
//	- only contains letters, numbers, and spaces
func SpaceSepToCamel(s string) string {
	validSpaceSepRe := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]*( [a-zA-Z0-9]+)*$`)
	if !validSpaceSepRe.MatchString(s) {
		log.Println("input is not a valid space separated string")
		return s
	}

	b := []byte(strings.TrimSpace(s))
	buffer := make([]byte, 0, len(s))

	tokens := bytes.Split(b, []byte(" "))
	for _, tok := range tokens {
		buffer = append(buffer, bytes.ToUpper(tok[:1])...)
		buffer = append(buffer, tok[1:]...)
	}

	return string(buffer)
}
