package easy

func runningSum(nums []int) []int {
	table := map[int]int{}
	result := make([]int, len(nums))

	for i, num := range nums {
		if i != 0 {
			num += table[i-1]
		}
		table[i] = num
		result[i] = num
	}
	return result
}
