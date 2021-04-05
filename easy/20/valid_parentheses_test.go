package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	cases := []struct {
		s      string
		expect bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"(([]){})", true},
		{"((", false},
		{"(([]){})", true},
	}

	for _, data := range cases {
		assert.Equal(t, isValid(data.s), data.expect)
	}
}
