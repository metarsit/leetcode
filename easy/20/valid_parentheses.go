package easy

func isValid(s string) bool {
	bStr := []byte(s)
	bracketSet := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	openBracketOrdered := []string{}

	for _, b := range bStr {
		currentBracket := string(b)
		if _, ok := bracketSet[currentBracket]; !ok {
			// Is Open Bracket
			openBracketOrdered = append(openBracketOrdered, currentBracket)
			continue
		}
		// If it starts with close bracket
		if len(openBracketOrdered) == 0 {
			return false
		}

		lastPos := len(openBracketOrdered) - 1
		openBracket := openBracketOrdered[lastPos]
		// Check if the open bracket matches what we need
		if currentOpenBracket, ok := bracketSet[currentBracket]; ok && currentOpenBracket != openBracket {
			return false
		}
		openBracketOrdered = openBracketOrdered[:lastPos]
	}
	return len(openBracketOrdered) == 0
}
