package easy

func shuffle(nums []int, n int) []int {
	sliceX := nums[:n]
	sliceY := nums[n:]

	output := make([]int, 0)
	for i := 0; i < n; i++ {
		output = append(output,
			sliceX[i], sliceY[i])
	}

	return output
}
