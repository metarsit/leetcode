package easy

func countGoodRectangles(rectangles [][]int) int {
	squaresLengths := make(map[int]int)
	largestSq := 0
	for _, rec := range rectangles {
		key := 0
		if rec[0] > rec[1] {
			squaresLengths[rec[1]] += 1
			key = rec[1]
		} else {
			squaresLengths[rec[0]] += 1
			key = rec[0]
		}
		if key > largestSq {
			largestSq = key
		}
	}
	return squaresLengths[largestSq]
}
