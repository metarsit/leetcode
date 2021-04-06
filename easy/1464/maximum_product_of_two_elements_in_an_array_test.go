package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProduct(t *testing.T) {
	cases := []struct {
		nums   []int
		expect int
	}{
		{[]int{3, 4, 5, 2}, 12},
		{[]int{1, 5, 4, 5}, 16},
		{[]int{3, 7}, 12},
	}

	for _, data := range cases {
		assert.Equal(t, maxProductWithoutLib(data.nums), data.expect)
		assert.Equal(t, maxProductWithLib(data.nums), data.expect)
	}
}

// MaxProduct with library is a lot slower as
// sort is a very expensive operation
func BenchmarkIsAnagram(b *testing.B) {
	sample := []int{3, 4, 5, 2}

	b.Run("Without Library", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			maxProductWithoutLib(sample)
		}
	})

	b.Run("With Sort Library", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			maxProductWithLib(sample)
		}
	})
}
