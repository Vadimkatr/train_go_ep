package factorial

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		n, want uint
	}{
		{0, 1},
		{1, 1},
		{3, 6},
		{5, 120},
		{10, 3628800},
	}
	for _, c := range tests {
		got := Factorial(c.n)
		if got != c.want {
			t.Errorf("Factorial(%d) == %d, want %d", c.n, got, c.want)
		}
	}
}
