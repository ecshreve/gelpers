package gelpers

// InIntSlice checks if the given slice contains the given value.
func InIntSlice(s []int, v int) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}
	return false
}
