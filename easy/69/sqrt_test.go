package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	cases := []struct {
		x      int
		expect int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{8, 2},
		{100, 10},
		{101, 10},
		{200, 14},
	}

	for _, data := range cases {
		assert.Equal(t, mySqrtWithMathLib(data.x), data.expect)
		assert.Equal(t, mySqrtWithoutLib(data.x), data.expect)
	}
}

// Sqrt without library is a lot slower as
// we have to compute all the possibility
// from 0 to the supposed answer
func BenchmarkMySqrt(b *testing.B) {
	sampleCase := 100

	b.Run("Without Library", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			mySqrtWithoutLib(sampleCase)
		}
	})

	b.Run("With Math Library", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			mySqrtWithMathLib(sampleCase)
		}
	})
}
