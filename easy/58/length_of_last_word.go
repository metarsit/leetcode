package easy

func lengthOfLastWord(s string) int {
	const space = 32
	b := []byte(s)
	lastWord := []byte{}

	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == space {
			if len(lastWord) == 0 {
				continue
			}
			break
		}
		lastWord = append(lastWord, b[i])
	}
	return len(lastWord)
}
