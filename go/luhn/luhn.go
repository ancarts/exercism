package luhn

import (
	"strconv"
	"unicode"
)

func Valid(id string) bool {
	idR := []rune(id)
	if len(id) <= 1 {
		return false
	}

	sum := 0
	n := 0

	for i := len(idR)-1; i >= 0; i-- {
		num, err := strconv.Atoi((string(idR[i])))
		if err != nil && !unicode.IsSpace(idR[i]) {
			return false
		}
		if unicode.IsSpace(idR[i]) {
			continue
		}
		if n%2 != 0 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
		n++
	}
	return sum%10 == 0&&n>1
}
