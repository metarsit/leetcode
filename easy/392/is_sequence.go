package easy

func isSubsequence(s string, t string) bool {
	sPos := 0
	tPos := 0

	for sPos < len(s) && tPos < len(t) {
		if s[sPos] == t[tPos] {
			sPos++
		}
		tPos++
	}
	return sPos == len(s)
}
