package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func shuffle(nums []int, n int) []int {
	sliceX := nums[:n]
	sliceY := nums[n:]

	output := make([]int, 0)
	for i := 0; i < n; i++ {
		output = append(output,
			sliceX[i], sliceY[i])
	}

	return output
}

func TestShuffle(t *testing.T) {
	cases := []struct {
		nums   []int
		n      int
		expect []int
	}{
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}

	for _, data := range cases {
		assert.Equal(t, shuffle(data.nums, data.n), data.expect)
	}
}
