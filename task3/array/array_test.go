package array

import (
	"testing"
)

func TestAverage(t *testing.T) {
	var tests = []struct{
		arr  []int
		wantMean float64
	}{
		{[]int{0}, 0},
		{[]int{1, 1, 1}, 1},
		{[]int{12, 0, 0}, 4.0},
		{[]int{1, 2, 3, 4, 5, 6}, 3.5},
	}
	for _, c := range tests {
		got := Average(c.arr)
		if got != c.wantMean {
			t.Errorf("Average(%v) == %f, want %f", c.arr, got, c.wantMean)
		}
	}
}
