package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
