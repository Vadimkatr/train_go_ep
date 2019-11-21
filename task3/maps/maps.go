package maps

import (
	"sort"
)

func PrintSorted(mp map[int]string) []string {
	sortedValues := make([]string, len(mp))
	keys := make([]int, 0, len(mp))
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i, key := range keys {
		sortedValues[i] = mp[key]
	}
	return sortedValues
}
