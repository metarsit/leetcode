package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1 <= n <= 1000
func TestSumZero(t *testing.T) {
	for i := 1; i <= 1000; i++ {
		result := 0
		for _, num := range sumZero(i) {
			result += num
		}

		assert.Equal(t, result, 0)
	}
}
