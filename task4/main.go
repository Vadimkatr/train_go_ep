package main

import (
	"fmt"
	"sort"
	"train_go_ep/task4/figure"
	"train_go_ep/task4/person"
	"train_go_ep/task4/utils"
)

func main() {
	ivanIvanovDate := utils.DateParse("2005-Aug-10")
	ivanIvanovSecondDate := utils.DateParse("2003-Aug-05")
	ivanBbbbDate := utils.DateParse("2005-Aug-10")
	artiomIvanovDate := utils.DateParse("2005-Aug-10")
	artiomPetrovDate := utils.DateParse("2010-Aug-10")
	arseniiAaaaDate := utils.DateParse("2005-Aug-10")
	dimaFedorovDate := utils.DateParse("2004-Jun-01")
	sahsaFedorovDate := utils.DateParse("2004-Jun-01")

	p := person.People{
		{"Ivan", "Ivanov", ivanIvanovDate},
		{"Ivan", "Ivanov", ivanIvanovSecondDate},
		{"Ivan", "Bddd", ivanBbbbDate},
		{"Artiom", "Ivanov", artiomIvanovDate},
		{"Artiom", "Petrov", artiomPetrovDate},
		{"Dima", "Fedorov", dimaFedorovDate},
		{"Sahsa", "Fedorov", sahsaFedorovDate},
		{"Arsenii", "Aaaa", arseniiAaaaDate},
	}

	fmt.Println(p)
	// 1) Ivan Ivanov 2005-08-10.
	// 2) Ivan Ivanov 2003-08-05.
	// 3) Ivan Bddd 2005-08-10.
	// 4) Artiom Ivanov 2005-08-10.
	// 5) Artiom Petrov 2010-08-10.
	// 6) Dima Fedorov 2004-06-01.
	// 7) Sahsa Fedorov 2004-06-01.
	// 8) Arsenii Aaaa 2005-08-10.
	sort.Sort(p)
	fmt.Println(p)
	// 1) Ivan Ivanov 2005-08-10.
	// 2) Ivan Ivanov 2003-08-05.
	// 3) Ivan Bddd 2005-08-10.
	// 4) Artiom Ivanov 2005-08-10.
	// 5) Artiom Petrov 2010-08-10.
	// 6) Dima Fedorov 2004-06-01.
	// 7) Sahsa Fedorov 2004-06-01.
	// 8) Arsenii Aaaa 2005-08-10.

	point := figure.Point{X: 1, Y: 1}
	var s figure.Figure = figure.Square{Start: point, A: 4, B: 5}
	var c figure.Figure = figure.Circle{Center: point, R: 3}

	fmt.Printf("Square area: %f; perimeter: %f\n", s.Area(), s.Perimeter())
	// Square area: 20.000000; perimeter: 18.000000
	fmt.Printf("Circle area: %f; perimeter: %f\n", c.Area(), c.Perimeter())
	// Circle area: 28.274334; perimeter: 18.849556
}
