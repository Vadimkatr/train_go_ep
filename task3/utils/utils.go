package utils

func EqualSliceOfInt64(firstSlice, secondSlice []int64) bool {
	if len(firstSlice) != len(secondSlice) {
		return false
	}
	for i := 0; i < len(firstSlice); i ++ {
		if firstSlice[i] != secondSlice[i] {
			return false
		}
	}
	return true
}

func EqualSliceOfStrings(firstSlice, secondSlice []string) bool {
	if len(firstSlice) != len(secondSlice) {
		return false
	}
	for i := 0; i < len(firstSlice); i ++ {
		if firstSlice[i] != secondSlice[i] {
			return false
		}
	}
	return true
}
