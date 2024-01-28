package lib

import (
	"fmt"
	"time"
)

func FormatDateAsString(date time.Time) string {
	day := padNumber(date.Day())
	month := padNumber(int(date.Month()))
	year := date.Year()

	return fmt.Sprintf("%s.%s.%d", day, month, year)
}

func padNumber(n int) string {
	if n < 10 {
		return fmt.Sprintf("0%d", n)
	}

	return fmt.Sprint(n)
}
