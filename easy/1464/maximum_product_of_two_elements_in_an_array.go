package easy

import "sort"

func maxProductWithoutLib(nums []int) int {
	firstLargest, secondLargest := 0, 0
	for _, currentNum := range nums {
		switch {
		case secondLargest > currentNum:
			continue
		case firstLargest < currentNum:
			firstLargest, secondLargest = currentNum, firstLargest
		default:
			secondLargest = currentNum
		}
	}
	return (firstLargest - 1) * (secondLargest - 1)
}

func maxProductWithLib(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}
