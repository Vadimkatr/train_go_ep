package main

import (
	"fmt"
	"train_go_ep/task3/arrays"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arrays.Average(arr))

	someStrings := []string{"a", "asd", "4ccc", "aa", "5aaaa"}
	fmt.Println(arrays.Max(someStrings))
	someStrings = []string{"a", "b"}
	fmt.Println(arrays.Max(someStrings))

}
