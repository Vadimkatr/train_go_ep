package converter

import (
	"fmt"
	"strconv"
)

func MyStrToIntUseAtoi(s string) (int, error) {
	return strconv.Atoi(s)
}

func MyStrToIntUseFmt(s string) (int, error) {
	var x int
	n, err := fmt.Sscanf(s, "%d\r", &x);
	if n != 1 || err != nil {
		return 0, err
	  }
	return x, nil
}
