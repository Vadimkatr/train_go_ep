package slices

import (
	"errors"
)

func Max(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("func Max(slice): slice is empty")
	}
	idxmax := 0
	maxlen := len(slice[0])

	for i, value := range slice {
		if len(value) > maxlen {
			maxlen = len(value)
			idxmax = i
		}
	}
	return slice[idxmax], nil
}

func Reverse(arr []int64) []int64 {
	reversearr := make([]int64, len(arr))
	size := len(arr)
	for i, value := range arr {
		reversearr[size-i-1] = value
	}
	return reversearr
}
