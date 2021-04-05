package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
