package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBits(t *testing.T) {
	cases := []struct {
		n      int
		expect int
	}{
		{2, 2},
		{3, 3},
	}

	for _, data := range cases {
		assert.Equal(t, climbStairs(data.n), data.expect)
	}
}
