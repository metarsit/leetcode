package easy

func validPalindrome(s string) bool {
	return isValid([]byte(s), 0, len(s)-1, false)
}

func isValid(b []byte, leftCursor, rightCursor int, removed bool) bool {
	for leftCursor <= rightCursor {
		switch {
		case b[leftCursor] == b[rightCursor]:
			leftCursor++
			rightCursor--
		case removed:
			return false
		default:
			return isValid(b, leftCursor+1, rightCursor, true) ||
				isValid(b, leftCursor, rightCursor-1, true)
		}
	}

	return true
}
