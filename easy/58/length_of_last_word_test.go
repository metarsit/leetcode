package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLastWord(s string) int {
	const space = 32
	b := []byte(s)
	lastWord := []byte{}

	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == space {
			if len(lastWord) == 0 {
				continue
			}
			break
		}
		lastWord = append(lastWord, b[i])
	}
	return len(lastWord)
}

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
