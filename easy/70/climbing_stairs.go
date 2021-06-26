package easy

func climbStairs(n int) int {
	prev, curr := 0, 1
	result := 1

	for i := 1; i <= n; i++ {
		result = prev + curr
		prev, curr = curr, result
	}

	return result
}
