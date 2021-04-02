package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isValid(s string) bool {
	bStr := []byte(s)
	bracketSet := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	openBracketOrdered := []string{}

	for _, b := range bStr {
		currentBracket := string(b)
		if _, ok := bracketSet[currentBracket]; !ok {
			// Is Open Bracket
			openBracketOrdered = append(openBracketOrdered, currentBracket)
			continue
		}
		// If it starts with close bracket
		if len(openBracketOrdered) == 0 {
			return false
		}

		lastPos := len(openBracketOrdered) - 1
		openBracket := openBracketOrdered[lastPos]
		// Check if the open bracket matches what we need
		if currentOpenBracket, ok := bracketSet[currentBracket]; ok && currentOpenBracket != openBracket {
			return false
		}
		openBracketOrdered = openBracketOrdered[:lastPos]
	}
	return len(openBracketOrdered) == 0
}

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
