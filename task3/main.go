package main

import (
	"fmt"
	"train_go_ep/task3/arrays"
	"train_go_ep/task3/maps"
	"train_go_ep/task3/slices"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arrays.Average(arr)) // 3.5

	someStrings := []string{"a", "asd", "4ccc", "aa", "5aaaa"}
	if result, err := slices.Max(someStrings); err == nil {
		fmt.Println(result) // 5aaaa
	}
	someStrings = []string{"a", "b"}
	if result, err := slices.Max(someStrings); err == nil {
		fmt.Println(result) // a
	}

	someSlice := []int64{1, 2, 3, 4, 5, 6}
	fmt.Println(slices.Reverse(someSlice)) // [6 5 4 3 2 1]

	myMap := map[int]string{
		2: "a",
		0: "b",
		1: "c",
	}
	fmt.Println(maps.PrintSorted(myMap)) // [b c a]
	myMapSecond := map[int]string{
		10:  "aa",
		0:   "bb",
		500: "cc",
	}
	fmt.Println(maps.PrintSorted(myMapSecond)) // [bb aa cc]
}
