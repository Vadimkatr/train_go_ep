package main

import (
	"fmt"
	"train_go_ep/task2/factorial"
)

func main() {
	var n uint = 5
	fmt.Printf("Factorial(%d) = %d\n", n, factorial.Factorial(n))
}
