package easy

import (
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	xi := strings.Split(s, " ")
	result := make([]string, len(xi))
	for _, word := range xi {
		lengthOfWord := len(word)
		pos, _ := strconv.Atoi(word[lengthOfWord-1:])
		result[pos-1] = word[:lengthOfWord-1]
	}
	return strings.Join(result, " ")
}
