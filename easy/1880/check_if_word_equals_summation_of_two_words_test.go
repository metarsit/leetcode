package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSumEqual(t *testing.T) {
	cases := []struct {
		firstWord  string
		secondWord string
		targetWord string
		expect     bool
	}{
		{"acb", "cba", "cdb", true},
		{"aaa", "a", "aab", false},
		{"aaa", "a", "aaaa", true},
	}

	for _, data := range cases {
		assert.Equal(t,
			isSumEqual(data.firstWord, data.secondWord, data.targetWord), data.expect)
	}
}
