package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fib(n int) int {
	var table = map[int]int{
		0: 0,
		1: 1,
	}

	if corr, ok := table[n]; ok {
		return corr
	}
	result := fib(n-1) + fib(n-2)
	table[n] = result
	return result
}

func TestFib(t *testing.T) {
	cases := []struct {
		n      int
		expect int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{12, 144},
		{13, 233},
		{14, 377},
		{15, 610},
		{16, 987},
		{17, 1597},
		{18, 2584},
		{19, 4181},
		{20, 6765},
		{21, 10946},
		{22, 17711},
		{23, 28657},
		{24, 46368},
		{25, 75025},
		{26, 121393},
		{27, 196418},
		{28, 317811},
		{29, 514229},
		{30, 832040},
	}

	for _, data := range cases {
		assert.Equal(t, fib(data.n), data.expect)
	}
}
