package utils

import (
	"fmt"
	"os"
	"time"
)

func DateParse(dateStr string) time.Time {
	date, err := time.Parse("2006-Jan-02", dateStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return date
}
