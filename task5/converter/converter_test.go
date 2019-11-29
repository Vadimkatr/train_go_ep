package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct{
	input   string
	output 	int
	isValid bool
}{
	{"", 0, false},
	{"0", 0, true},
	{"-0", 0, true},
	{"1", 1, true},
	{"-1", -1, true},
	{"12345", 12345, true},
	{"-12345", -12345, true},
	{"012345", 12345, true},
	{"-012345", -12345, true},
	{"98765432100", 98765432100, true},
	{"-98765432100", -98765432100, true},
	{"9223372036854775807", 1<<63 - 1, true},
	{"-9223372036854775807", -(1<<63 - 1), true},
	{"9223372036854775808", 1<<63 - 1, false},
	{"-9223372036854775808", -1 << 63, true},
	{"9223372036854775809", 1<<63 - 1, false},
	{"-9223372036854775809", -1 << 63, false},
	{"-1_2_3_4_5", 0, false}, // base=10 so no underscores allowed
	{"-_12345", 0, false},
	{"_12345", 0, false},
	{"1__2345", 0, false},
	{"12345_", 0, false},
}

func TestMyStrToIntUseAtoi(t *testing.T) {
	for _, tc := range testCases {
		if tc.isValid {
			n, err := MyStrToIntUseAtoi(tc.input)
			assert.Equal(t, n, tc.output)
			assert.Equal(t, err, nil)
		} else {
			_, err := MyStrToIntUseAtoi(tc.input)
			assert.NotEqual(t, err, nil)
		}
	}
}

func TestMyStrToIntUseFmt(t *testing.T) {
	for _, tc := range testCases {
		if tc.isValid {
			n, err := MyStrToIntUseFmt(tc.input)
			assert.Equal(t, n, tc.output)
			assert.Equal(t, err, nil)
		} else {
			_, err := MyStrToIntUseFmt(tc.input)
			assert.NotEqual(t, err, nil)
		}
	}
}

func BenchmarkMyStrToIntUseAtoi(b *testing.B) {
    for i := 0; i < b.N; i++ {
		_, _ = MyStrToIntUseAtoi(string(i))
    }
}

func BenchmarkMyStrToIntUseFmt(b *testing.B) {
    for i := 0; i < b.N; i++ {
		_, _ = MyStrToIntUseFmt(string(i))
    }
}
