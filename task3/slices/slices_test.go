package slices

import (
	"reflect"
	"testing"
)

func TestMaxMainCases(t *testing.T) {
	var tests = []struct {
		arr       []string
		maxString string
	}{
		{[]string{"a", "bb", "ccc"}, "ccc"},
		{[]string{"a", "b", "c"}, "a"},
		{[]string{"aaa", "b"}, "aaa"},
		{[]string{""}, ""},
	}
	for _, c := range tests {
		got, _ := Max(c.arr)
		if got != c.maxString {
			t.Errorf("Max(%v) == %s, want %s", c.arr, got, c.maxString)
		}
	}
}

func TestMaxEmptyArray(t *testing.T) {
	if _, err := Max([]string{}); err == nil {
		t.Errorf("Error: %s", err)
	}
}

func TestReverse(t *testing.T) {
	var tests = []struct {
		arr, reverseArr []int64
	}{
		{[]int64{}, []int64{}},
		{[]int64{1, 2, 3, 4, 5}, []int64{5, 4, 3, 2, 1}},
		{[]int64{1, 1, 1}, []int64{1, 1, 1}},
	}
	for _, c := range tests {
		gotArr := Reverse(c.arr)
		if !reflect.DeepEqual(gotArr, c.reverseArr) {
			t.Errorf("Reverse(%v) == %v, want %v", c.arr, gotArr, c.reverseArr)
		}
	}
}
