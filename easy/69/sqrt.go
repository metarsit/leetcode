package easy

import (
	"math"
)

func mySqrtWithMathLib(x int) int {
	return int(math.Sqrt(float64(x)))
}

func mySqrtWithoutLib(x int) int {
	num := 0
	for (num * num) <= x {
		num++
	}

	return num - 1
}
