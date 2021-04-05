package easy

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBits(t *testing.T) {
	cases := []struct {
		num    string
		expect uint32
	}{
		{"00000010100101000001111010011100", 964176192},
		{"11111111111111111111111111111101", 3221225471},
	}

	for i, data := range cases {
		num, err := strconv.ParseUint(data.num, 2, 32)
		if err != nil {
			t.Fatalf("Test case #%d (%s) error: %s", i, data.num, err)
		}
		assert.Equal(t, reverseBits(uint32(num)), data.expect)
	}
}
