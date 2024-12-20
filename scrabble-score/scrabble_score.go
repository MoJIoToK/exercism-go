package scrabble

import (
	"strings"
	// "unicode"
)

// Sacrabble is letters points table
// var scrabble = map[rune]int{
var scrabble = map[string]int{
	// 'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	// 'D': 2, 'G': 2,
	// 'B': 3, 'C': 3, 'M': 3, 'P': 3,
	// 'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	// 'K': 5,
	// 'J': 8, 'X': 8,
	// 'Q': 10, 'Z': 10,
	"A": 1, "E": 1, "I": 1, "O": 1, "U": 1, "L": 1, "N": 1, "R": 1, "S": 1, "T": 1,
	"D": 2, "G": 2,
	"B": 3, "C": 3, "M": 3, "P": 3,
	"F": 4, "H": 4, "V": 4, "W": 4, "Y": 4,
	"K": 5,
	"J": 8, "X": 8,
	"Q": 10, "Z": 10,
}

// Score calculates points according to letters in the points table.
func Score(word string) int {

	var scrabbleScore int
	for _, val := range word {
		// scrabbleScore += scrabble[unicode.ToUpper(val)]
		scrabbleScore += scrabble[strings.ToUpper(string(val))]
	}
	return scrabbleScore
}
