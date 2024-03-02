package isogram

import "strings"

func IsIsogram(word string) bool {
	toLoweWord := strings.ToLower(word)
	for _, r := range toLoweWord {
		if strings.Count(toLoweWord, string(r)) > 1 && string(r) != " " && string(r) != "-" {
			return false
		}
	}
	return true
}
