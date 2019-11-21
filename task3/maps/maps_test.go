package maps

import (
	"testing"
	"train_go_ep/task3/utils"
)

func TestPrintSortedMainCases(t *testing.T) {
	var tests = []struct {
		mp          map[int]string
		wantStrings []string
	}{
		{map[int]string{2: "c", 1: "b", 0: "a"}, []string{"a", "b", "c"}},
		{map[int]string{0: "c", 1: "a", 2: "b"}, []string{"c", "a", "b"}},
		{map[int]string{10: "b", -10: "a"}, []string{"a", "b"}},
		{map[int]string{100: "a"}, []string{"a"}},
		{map[int]string{}, []string{}},
	}
	for _, c := range tests {
		got := PrintSorted(c.mp)
		if !utils.EqualSliceOfStrings(got, c.wantStrings) {
			t.Errorf("Average(%v) == %v, want %v", c.mp, got, c.wantStrings)
		}
	}
}
