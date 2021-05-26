package easy

func twoSum(nums []int, target int) []int {
	lookupTable := make(map[int]int)
	for i, num := range nums {
		if pos, ok := lookupTable[num]; !ok {
			lookupTable[target-num] = i
		} else {
			return []int{pos, i}
		}
	}
	return nil
}
