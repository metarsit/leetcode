package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCenter(t *testing.T) {
	cases := []struct {
		edges  [][]int
		expect int
	}{
		{[][]int{{1, 2}, {2, 3}, {4, 2}}, 2},
		{[][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}, 1},
	}

	for _, data := range cases {
		assert.Equal(t, findCenter(data.edges), data.expect)
	}
}
