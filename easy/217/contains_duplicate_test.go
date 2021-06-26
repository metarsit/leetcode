package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBits(t *testing.T) {
	cases := []struct {
		nums   []int
		expect bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, data := range cases {
		assert.Equal(t, containsDuplicate(data.nums), data.expect)
	}
}
