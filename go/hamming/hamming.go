package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		errLenghtIsNotEqual := errors.New("lenght is not equal")
		return 0, errLenghtIsNotEqual
	}

	count := 0
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
