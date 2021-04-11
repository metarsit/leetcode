package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		nums   []int
		expect int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{0}, 1},
	}

	for _, data := range cases {
		assert.Equal(t, missingNumber(data.nums), data.expect)
	}
}
