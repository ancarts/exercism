package collatzconjecture

import (
	"errors"
	"fmt"
)

func CollatzConjecture(n int) (int, error) {

	if n <= 0 {
		var ErrNotAPositiveinteger = errors.New(fmt.Sprint(n) + "is not a positive integer")
		return 0, ErrNotAPositiveinteger
	}
	var count int
	for n != 1 {
		
		if n%2 != 0 {
		n = n*3 + 1
		 } else {
		n = n / 2
		 }
		count++
	}

	return count, nil
}
