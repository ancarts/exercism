package scrabble

import "strings"

func Score(word string) int {
	score := 0
	lettersTovalue := map[string]int{
		"AEIOULNRST": 1,
		"DG":         2,
		"BCMP":       3,
		"FHVWY":      4,
		"K":          5,
		"JX":         8,
		"QZ":         10,
	}

	wordUpper := strings.ToUpper(word)
	for _, v := range wordUpper {
		for k, j := range lettersTovalue {
			if strings.ContainsRune(k, v) {
				score += j
			}

		}
	}
	return score
}
