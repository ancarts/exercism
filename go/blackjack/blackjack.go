package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.

func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case card == "ten" || card == "jack" || card == "queen" || card == "king":
		return 10
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	value1 := ParseCard(card1)
	value2 := ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case value1+value2 == 21 && dealerValue < 10:
		return "W"
	case (value1+value2 == 21 && dealerValue >= 10) || (value1+value2 >= 17 && value1+value2 < 21) ||
		(value1+value2 >= 12 && value1+value2 <= 16 && dealerValue < 7):
		return "S"
	}
	return "H"
}
