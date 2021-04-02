package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func restoreString(s string, indices []int) string {
	xi := make([]byte, len(indices))
	b := []byte(s)
	for i, pos := range indices {
		xi[pos] = b[i]
	}
	return string(xi)
}

func TestRestoreString(t *testing.T) {
	cases := []struct {
		s       string
		indices []int
		expect  string
	}{
		{"codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}, "leetcode"},
		{"abc", []int{0, 1, 2}, "abc"},
		{"aiohn", []int{3, 1, 4, 2, 0}, "nihao"},
		{"aaiougrt", []int{4, 0, 2, 6, 7, 3, 1, 5}, "arigatou"},
		{"art", []int{1, 0, 2}, "rat"},
	}

	for _, data := range cases {
		assert.Equal(t, restoreString(data.s, data.indices), data.expect)
	}
}
