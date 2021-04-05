package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	cases := []struct {
		s      string
		expect int
	}{
		{"Hello World", 5},
		{" ", 0},
		{"a  ", 1},
		{"  a  ", 1},
		{" Some", 4},
	}

	for _, data := range cases {
		assert.Equal(t, lengthOfLastWord(data.s), data.expect)
	}
}
