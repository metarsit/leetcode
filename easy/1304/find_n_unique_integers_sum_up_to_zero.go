package easy

func sumZero(n int) []int {
	output := make([]int, n)
	if n == 1 {
		return output
	}

	current := 0
	for i := range output {
		switch {
		case i == (n - 1):
			if current != 0 {
				output[i] = -current
			}
		case i%2 == 0:
			output[i] = n + i
		default:
			output[i] = -(n + i - 1)
		}
		current += output[i]
	}
	return output
}
