package main

import (
	"fmt"
	"train_go_ep/task2/factorial"
	"train_go_ep/task2/point"
)

func main() {
	var n uint = 5
	fmt.Printf("Factorial(%d) = %d\n", n, factorial.Factorial(n))

	s := point.Square{point.Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
