package easy

func findCenter(edges [][]int) int {
	a0, b0 := edges[0][0], edges[0][1]
	a1, b1 := edges[1][0], edges[1][1]

	var center int
	switch {
	case (a0 == a1 || a0 == b1):
		center = a0
	case (b0 == a1 || b0 == b1):
		center = b0
	}

	return center
}
