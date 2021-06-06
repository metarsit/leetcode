package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortSentence(t *testing.T) {
	cases := []struct {
		s      string
		expect string
	}{
		{"is2 sentence4 This1 a3", "This is a sentence"},
	}

	for _, data := range cases {
		assert.Equal(t, sortSentence(data.s), data.expect)
	}
}
