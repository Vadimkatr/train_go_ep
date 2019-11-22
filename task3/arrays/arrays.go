package arrays

func Average(arr []int) float64 {
	total := 0
	for _, value := range arr {
		total += value
	}
	return float64(total) / float64(len(arr))
}
