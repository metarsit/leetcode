package easy

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = i
	}
	return false
}
