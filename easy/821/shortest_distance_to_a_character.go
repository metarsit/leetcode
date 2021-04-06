package easy

func shortestToChar(s string, c byte) []int {
	b := []byte(s)
	pos := make([]int, 0)

	for i, char := range b {
		if char == c {
			pos = append(pos, i)
		}
	}

	maxPosIdx := len(pos) - 1
	lengthOfB := len(b)
	output := make([]int, lengthOfB)
	j := 0

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for i := 0; i < lengthOfB; i++ {
		v := abs(i - pos[j])
		if j < maxPosIdx {
			altDist := abs(i - pos[j+1])
			if altDist < v {
				v = altDist
				j++
			}
		}
		output[i] = v
	}
	return output
}
