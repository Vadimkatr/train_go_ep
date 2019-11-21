package arrays

func Average(arr []int) float64 {
	total := 0
	for _, value := range arr {
		total += value
	}
	return float64(total) / float64(len(arr))
}


func Max(slice []string) string {
	if len(slice) == 0 {
		panic("func Max(slice): slice is empty")
	}
	idxmax := 0
	maxlen := len(slice[0])

	for i, value := range slice {
		if len(value) > maxlen {
			maxlen = len(value)
			idxmax = i
		}
	}
	return slice[idxmax]
}
