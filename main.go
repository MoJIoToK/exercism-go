package main

import (
	"fmt"
	"strings"
)

func main() {

	hash := map[int]int{9: 10, 99: 20, 999: 30}
	for k, v := range hash {
		fmt.Println(k, v)
	}

}

func IsIsogram(word string) bool {

	seen := make(map[rune]int)
	word = strings.Replace(word, "-", "", -1)
	//println("1", word)
	word = strings.Replace(word, " ", "", -1)
	//println("2", word)
	word = strings.ToLower(word)
	//println("3", word)

	for _, val := range word {
		seen[val]++
		if seen[val] > 1 {
			return false
		}
	}

	return true
}

//func IsIsogram(word string) bool {
//	// 'seen' is a slice of type bool
//	//seen := make([]bool, 26)
//	//for _, letter := range word {
//	//	// 'letter' is a rune
//	//	if letter == '-' || letter == ' ' {
//	//		continue
//	//	}
//	//	// make letter uppercase
//	//	if !(letter >= 'A' && letter <= 'Z') {
//	//		letter = letter - ('a' - 'A')
//	//	}
//	//	if seen[letter-'A'] {
//	//		//println(letter-'A', string(letter-'A'))
//	//		println(seen[letter-'A'])
//	//		return false
//	//	}
//	//	seen[letter-'A'] = true
//	//	println(letter-'A', string(letter-'A'))
//	//}
//
//	//return true
//}
