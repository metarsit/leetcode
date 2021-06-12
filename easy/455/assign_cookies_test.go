package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindContentChildren(t *testing.T) {
	cases := []struct {
		g      []int
		s      []int
		expect int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
	}

	for _, data := range cases {
		assert.Equal(t, findContentChildren(data.g, data.s), data.expect)
	}
}
