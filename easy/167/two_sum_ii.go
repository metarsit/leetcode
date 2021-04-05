package easy

func twoSum(numbers []int, target int) []int {
	table := make(map[int]int, len(numbers))

	for j, num := range numbers {
		if i, ok := table[num]; ok {
			return []int{i + 1, j + 1}
		}
		table[target-num] = j
	}
	return nil
}
