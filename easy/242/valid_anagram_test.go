package leetcode

import (
	"bytes"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isAnagramWithoutLibrary(s, t string) bool {
	sb := []byte(s)
	tb := []byte(t)

	if len(sb) != len(tb) {
		return false
	}

	table := make(map[byte]int)
	for i := 0; i <= (len(sb) - 1); i++ {
		currS := sb[i]
		if _, ok := table[currS]; !ok {
			table[currS] = 1
		} else {
			table[currS] += 1
		}

		currT := tb[i]
		if _, ok := table[currT]; !ok {
			table[currT] = -1
		} else {
			table[currT] -= 1
		}
	}

	for _, occ := range table {
		if occ != 0 {
			return false
		}
	}
	return true
}

func isAnagramWithLibrary(s, t string) bool {
	sb := []byte(s)
	tb := []byte(t)

	sort.SliceStable(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})
	sort.SliceStable(tb, func(i, j int) bool {
		return tb[i] < tb[j]
	})

	return bytes.Equal(tb, sb)
}

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
