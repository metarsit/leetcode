package easy

import (
	"strconv"
)

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	var charTable = map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
		"i": 8,
		"j": 9,
	}
	wordToNum := func(word string) int {
		var num string
		for _, char := range word {
			num += strconv.Itoa(charTable[string(char)])
		}

		result, _ := strconv.Atoi(string(num))
		return result
	}

	firstNum := wordToNum(firstWord)
	secondNum := wordToNum(secondWord)
	targetNum := wordToNum(targetWord)
	return firstNum+secondNum == targetNum
}
