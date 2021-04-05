package easy

import (
	"time"
)

// dayOfTheWeekWithLibrary uses time library to parse the data
// and return the desired date string
func dayOfTheWeekWithLibrary(day, month, year int) string {
	return time.
		Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).
		Weekday().
		String()
}
