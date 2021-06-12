package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIdenticalPairs(t *testing.T) {
	cases := []struct {
		nums   []int
		expect int
	}{
		{[]int{1, 2, 3, 1, 1, 3}, 4},
		{[]int{1, 1, 1, 1}, 6},
		{[]int{1, 2, 3}, 0},
	}

	for _, data := range cases {
		assert.Equal(t, numIdenticalPairs(data.nums), data.expect)
	}
}
