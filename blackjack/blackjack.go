package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var score int
	switch card {
	case "ace":
		score = 11
	case "two":
		score = 2
	case "three":
		score = 3
	case "four":
		score = 4
	case "five":
		score = 5
	case "six":
		score = 6
	case "seven":
		score = 7
	case "eight":
		score = 8
	case "nine":
		score = 9
	case "ten":
		score = 10
	case "jack":
		score = 10
	case "queen":
		score = 10
	case "king":
		score = 10
	default:
		score = 0
	}
	return score
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var operation string = ""
	var score int = ParseCard(card1) + ParseCard(card2)
	switch {
	case card1 == "ace" && card2 == "ace":
		operation += "P"
	case score >= 17 && score <= 20:
		operation += "S"
	case score <= 11:
		operation += "H"
	case score >= 12 && score <= 16 && ParseCard(dealerCard) < 7:
		operation += "S"
	case score >= 12 && score <= 16 && ParseCard(dealerCard) >= 7:
		operation += "H"
	case score == 21 && ParseCard(dealerCard) < 10:
		operation += "W"
	case score == 21 && ParseCard(dealerCard) > 10:
		operation += "S"
	default:
		operation += "S"
	}
	return operation
}
