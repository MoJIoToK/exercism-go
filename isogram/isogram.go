package isogram

import "strings"

// IsIsogram returns true if there are no duplicate letters in input string.
func IsIsogram(word string) bool {
	word = strings.Replace(word, "-", "", -1)
	word = strings.Replace(word, " ", "", -1)
	word = strings.ToLower(word)
	seen := make(map[rune]int)
	for _, val := range word {
		seen[val]++
		if seen[val] > 1 {
			return false
		}
	}

	return true

}
