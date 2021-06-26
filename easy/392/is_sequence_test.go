package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSequence(t *testing.T) {
	cases := []struct {
		s      string
		t      string
		expect bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}

	for _, data := range cases {
		assert.Equal(t, isSubsequence(data.s, data.t), data.expect)
	}
}
