package easy

func numIdenticalPairs(nums []int) int {
	lengthOfNums := len(nums) - 1
	count := 0
	i := 0
	j := lengthOfNums

	for i < lengthOfNums {
		if i >= j {
			i++
			j = lengthOfNums
			continue
		}
		if nums[i] == nums[j] {
			count++
		}
		j--
	}
	return count
}
