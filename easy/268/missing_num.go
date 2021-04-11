package easy

func missingNumber(nums []int) int {
	n := len(nums)
	triNumber := n * (n + 1) / 2 //nolint:gomnd // Formulae

	for _, num := range nums {
		triNumber -= num
	}

	return triNumber
}
