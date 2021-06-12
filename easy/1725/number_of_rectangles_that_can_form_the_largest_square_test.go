package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountGoodRectangles(t *testing.T) {
	cases := []struct {
		rectangles [][]int
		expect     int
	}{
		{[][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}, 3},
		{[][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}, 3},
	}

	for _, data := range cases {
		assert.Equal(t, countGoodRectangles(data.rectangles), data.expect)
	}
}
