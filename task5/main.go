package main

import (
	"fmt"

	"train_go_ep/task5/converter"
	"train_go_ep/task5/thirdtask"
)

func main() {
	// task 1: implement myStrToInt func with tests and coverage;
	// see func converter.MyStrToIntUseAtoi,
	// converter/converter_test.go for tests;
	// and converter/coverout.html for coverage;
	text := "123456789"
	n, err := converter.MyStrToIntUseAtoi(text)
	if err != nil {
		fmt.Printf("MyStrToIntUseAtoi Error %v\n", err)
	} else {
		fmt.Printf("MyStrToIntUseAtoi Number %d\n", n)
	}

	// task 2: implement another myStrToInt func that pass tests of first;
	// see func converter.MyStrToIntUseFmt,
	// converter/converter_test.go for tests;
	// and converter/tests_results.txt for benchmarks;
	n, err = converter.MyStrToIntUseFmt(text)
	if err != nil {
		fmt.Printf("MyStrToIntUseFmt Error %v\n", err)
	} else {
		fmt.Printf("MyStrToIntUseFmt Number %d\n", n)
	}

	// task 3: fix error in code
	thirdtask.MainOfTherdTask()
}
