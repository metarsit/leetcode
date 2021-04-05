package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
