package main

import (
	"fmt"
	"train_go_ep/task3/arrays"
	"train_go_ep/task3/slices"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arrays.Average(arr))  // 3.5

	someStrings := []string{"a", "asd", "4ccc", "aa", "5aaaa"}
	fmt.Println(slices.Max(someStrings))  // 5aaaa
	someStrings = []string{"a", "b"}
	fmt.Println(slices.Max(someStrings))  // a

	someSlice := []int64{1, 2, 3, 4, 5, 6}
	fmt.Println(slices.Reverse(someSlice))  // [6 5 4 3 2 1]
}
