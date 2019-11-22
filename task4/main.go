package main

import (
	"fmt"
	"os"
	"sort"
	"time"
	"train_go_ep/task4/person"
)

type People []person.Person

func main() {
	ivanIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p := People{
		{"Ivan", "Ivanov", ivanIvanovDate},
	}

	fmt.Println(sort.Sort(p))
}
