package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func runningSum(nums []int) []int {
	table := map[int]int{}
	result := make([]int, len(nums))

	for i, num := range nums {
		if i != 0 {
			num += table[i-1]
		}
		table[i] = num
		result[i] = num
	}
	return result
}

func TestRunningSum(t *testing.T) {
	cases := []struct {
		nums   []int
		expect []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}

	for _, data := range cases {
		assert.Equal(t, runningSum(data.nums), data.expect)
	}
}
