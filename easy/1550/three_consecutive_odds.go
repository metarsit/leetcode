package easy

func threeConsecutiveOdds(arr []int) bool {
	const (
		evenable       = 2
		criteriaConsec = 3
	)

	consecOdd := 0

	for _, num := range arr {
		if (num % evenable) == 0 {
			consecOdd = 0
			continue
		}

		consecOdd++
		if consecOdd == criteriaConsec {
			return true
		}
	}
	return false
}
