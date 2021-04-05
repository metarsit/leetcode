package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPalindrome(t *testing.T) {
	cases := []struct {
		s      string
		expect bool
	}{
		{"aba", true},
		{"abc", false},
		{"abca", true},
		{"deeee", true},
		{"eedede", true},
		{"ebcbbececabbacecbbcbe", true},
		{"aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga", true},
	}

	for _, data := range cases {
		assert.Equal(t, validPalindrome(data.s), data.expect)
	}
}
