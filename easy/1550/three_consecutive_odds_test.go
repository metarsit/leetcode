package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func threeConsecutiveOdds(arr []int) bool {
	consecOdd := 0

	for _, num := range arr {
		if (num % 2) == 0 {
			consecOdd = 0
			continue
		}

		consecOdd++
		if consecOdd == 3 {
			return true
		}
	}
	return false
}

func TestThreeConsecutiveOdds(t *testing.T) {
	cases := []struct {
		arr    []int
		expect bool
	}{
		{[]int{2, 6, 4, 1}, false},
		{[]int{1, 2, 34, 3, 4, 5, 7, 23, 12}, true},
	}

	for _, data := range cases {
		assert.Equal(t, threeConsecutiveOdds(data.arr), data.expect)
	}
}
