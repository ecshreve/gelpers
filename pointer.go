package gelpers

// StringPtr conveniently converts a string literal into a pointer to string.
func StringPtr(str string) *string {
	return &str
}

// StringVal conveniently converts a string literal for the value of the given
// pointer to string, if the given pointer is nil it returns an empty string.
func StringVal(strPtr *string) string {
	if strPtr == nil {
		return ""
	}
	return *strPtr
}
