package maps

import (
	"sort"
)

func PrintSorted(mp map[int]string) []string {
	sortedValues := make([]string, len(mp))
	keys := make([]int, len(mp), len(mp))
	i := 0
	for k, _ := range mp {
		keys[i] = k
		i ++
	}
	sort.Ints(keys)
	for i, key := range keys {
		sortedValues[i] = mp[key]
	}
	return sortedValues
}
