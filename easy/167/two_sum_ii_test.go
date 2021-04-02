package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(numbers []int, target int) []int {
	table := make(map[int]int, len(numbers))

	for j, num := range numbers {
		if i, ok := table[num]; ok {
			return []int{i + 1, j + 1}
		}
		table[target-num] = j
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	cases := []struct {
		numbers []int
		target  int
		expect  []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	for _, data := range cases {
		assert.Equal(t,
			twoSum(data.numbers, data.target), data.expect)
	}
}
