package person

import (
	"reflect"
	"sort"
	"testing"
	"train_go_ep/task4/utils"
)

func TestPeopleSort(t *testing.T) {
	AaaAaaDate := utils.DateParse("2005-Aug-10")
	AaaAaaSecondDate := utils.DateParse("2003-Aug-10")
	AaaBbbDate := utils.DateParse("2005-Aug-10")
	BbbAaaDate := utils.DateParse("2006-Aug-10")
	var tests = []struct {
		arr, sortArr People
	}{
		{
			People{
				{"Aaa", "Aaa", AaaAaaDate}, // [Aaa Aaa 2005-Aug-10]
			},
			People{
				{"Aaa", "Aaa", AaaAaaDate}, // [Aaa Aaa 2005-Aug-10]
			},
		},
		{
			People{
				{"Aaa", "Bbb", AaaBbbDate}, // [Aaa Bbb 2005-Aug-10]
				{"Aaa", "Aaa", AaaAaaDate}, // [Aaa Aaa 2005-Aug-10]
			},
			People{
				{"Aaa", "Aaa", AaaBbbDate}, // [Aaa Aaa 2005-Aug-10]
				{"Aaa", "Bbb", AaaBbbDate}, // [Aaa Bbb 2005-Aug-10]

			},
		},
		{
			People{
				{"Bbb", "Aaa", BbbAaaDate}, // [Bbb Aaa 2006-Aug-10]
				{"Aaa", "Aaa", AaaAaaDate}, // [Aaa Aaa 2005-Aug-10]
			},
			People{
				{"Bbb", "Aaa", BbbAaaDate}, // [Bbb Aaa 2006-Aug-10]
				{"Aaa", "Aaa", AaaBbbDate}, // [Aaa Aaa 2005-Aug-10]

			},
		},
		{
			People{
				{"Aaa", "Aaa", AaaAaaSecondDate}, // [Aaa Aaa 2003-Aug-10]
				{"Aaa", "Aaa", AaaAaaDate}, // [Aaa Aaa 2005-Aug-10]
				{"Aaa", "Bbb", AaaBbbDate}, // [Aaa Bbb 2005-Aug-10]
			},
			People{
				{"Aaa", "Aaa", AaaAaaDate}, // [Aaa Aaa 2005-Aug-10]
				{"Aaa", "Bbb", AaaBbbDate}, // [Aaa Bbb 2005-Aug-10]
				{"Aaa", "Aaa", AaaAaaSecondDate}, // [Aaa Aaa 2003-Aug-10]
			},
		},
	}
	for _, c := range tests {
		sort.Sort(c.arr)
		if !reflect.DeepEqual(c.arr, c.sortArr) {
			t.Errorf("Arr \n%v after sort must be \n%v", c.arr, c.sortArr)
		}
	}
}
