package easy

func reverseString(s []byte) {
	rightCur := len(s) - 1
	leftCur := 0

	for leftCur <= rightCur {
		s[rightCur], s[leftCur] =
			s[leftCur], s[rightCur]

		rightCur--
		leftCur++
	}
}
