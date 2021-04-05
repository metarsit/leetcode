package easy

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
