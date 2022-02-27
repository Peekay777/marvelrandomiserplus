package utils

// Find val in slice and return index and boolean
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// Is val in slice
func IsInSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// PrintArr returns a string with each element seperated by comma
func PrintArr(arr []string) string {
	var line string = "["
	for i, element := range arr {
		line = line + element
		if i != len(arr)-1 {
			line = line + ", "
		}
	}
	return line + "]"
}
