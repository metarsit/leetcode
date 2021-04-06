package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestToChar(t *testing.T) {
	cases := []struct {
		s      string
		c      byte
		expect []int
	}{
		{"loveleetcode", []byte("e")[0], []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}},
		{"aaab", []byte("b")[0], []int{3, 2, 1, 0}},
		{"baaa", []byte("b")[0], []int{0, 1, 2, 3}},
	}

	for _, data := range cases {
		assert.Equal(t, shortestToChar(data.s, data.c), data.expect)
	}
}
