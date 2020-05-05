package exercises

// CopySlice copy strings
func CopySlice(strings []string) ([]string, []string) {
	var result []string
	result = make([]string, len(strings))

	copy(result, strings)

	result[0] = "Hello World"

	return strings, result
}
