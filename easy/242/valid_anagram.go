package easy

import (
	"bytes"
	"sort"
)

func isAnagramWithoutLibrary(s, t string) bool {
	sb := []byte(s)
	tb := []byte(t)

	if len(sb) != len(tb) {
		return false
	}

	table := make(map[byte]int)
	for i := 0; i <= (len(sb) - 1); i++ {
		currS := sb[i]
		if _, ok := table[currS]; !ok {
			table[currS] = 1
		} else {
			table[currS]++
		}

		currT := tb[i]
		if _, ok := table[currT]; !ok {
			table[currT] = -1
		} else {
			table[currT]--
		}
	}

	for _, occ := range table {
		if occ != 0 {
			return false
		}
	}
	return true
}

func isAnagramWithLibrary(s, t string) bool {
	sb := []byte(s)
	tb := []byte(t)

	sort.SliceStable(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})
	sort.SliceStable(tb, func(i, j int) bool {
		return tb[i] < tb[j]
	})

	return bytes.Equal(tb, sb)
}
