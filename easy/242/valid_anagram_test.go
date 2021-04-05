package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		s      string
		t      string
		expect bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"ab", "a", false},
		{"abcde", "edcba", true},
		{"åß∂ƒ©ƒ∂ßå", "∂åß∂ßåƒ©ƒ", true},
	}

	for _, data := range cases {
		assert.Equal(t, isAnagramWithoutLibrary(data.s, data.t), data.expect)
		assert.Equal(t, isAnagramWithLibrary(data.s, data.t), data.expect)
	}
}

// IsAnagram with library is a lot slower as
// sort is a very expensive operation
func BenchmarkIsAnagram(b *testing.B) {
	s := "anagram"
	t := "nagaram"

	b.Run("Without Library", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			isAnagramWithoutLibrary(s, t)
		}
	})

	b.Run("With Sort Library", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			isAnagramWithLibrary(s, t)
		}
	})
}
