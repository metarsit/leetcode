package leetcode

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// dayOfTheWeekWithLibrary uses time library to parse the data
// and return the desired date string
func dayOfTheWeekWithLibrary(day, month, year int) string {
	return time.
		Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).
		Weekday().
		String()
}

func TestDayOfTheWeek(t *testing.T) {
	cases := []struct {
		day    int
		month  int
		year   int
		expect string
	}{
		{31, 8, 2019, "Saturday"},
		{18, 7, 1999, "Sunday"},
		{15, 8, 1993, "Sunday"},
	}

	for _, data := range cases {
		assert.Equal(t, dayOfTheWeekWithLibrary(
			data.day, data.month, data.year), data.expect)
	}
}
