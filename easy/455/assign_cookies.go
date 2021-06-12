package easy

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i := 0
	j := 0
	count := 0

	lengthOfG := len(g)
	lengthOfS := len(s)

	for i < lengthOfG && j < lengthOfS {
		if s[j] >= g[i] {
			i++
			count++
		}

		j++
	}

	return count
}
