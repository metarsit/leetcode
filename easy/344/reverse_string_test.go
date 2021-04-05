package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		s      *[]byte
		expect []byte
	}{
		{&[]byte{5, 4, 3, 2, 1}, []byte{1, 2, 3, 4, 5}},
		{&[]byte{100, 4, 3, 2, 100}, []byte{100, 2, 3, 4, 100}},
		{&[]byte{5, 1}, []byte{1, 5}},
		{&[]byte{1}, []byte{1}},
		{&[]byte{1, 0, 0, 0, 0, 0, 0, 1}, []byte{1, 0, 0, 0, 0, 0, 0, 1}},
	}

	for _, data := range cases {
		reverseString(*data.s)
		assert.Equal(t, *data.s, data.expect)
	}
}
