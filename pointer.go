package gelpers

// StringPtr conveniently converts a string literal into a pointer to string.
func StringPtr(str string) *string {
	return &str
}

// StringVal conveniently converts a pointer to string to a string literal, if
// the given pointer is nil it returns an empty string.
func StringVal(strPtr *string) string {
	if strPtr == nil {
		return ""
	}
	return *strPtr
}
