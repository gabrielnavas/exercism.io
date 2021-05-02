package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("size of 'a' is different from 'b'")
	}

	lenHamming := len(a)
	errorsCount := 0
	for index := 0; index < lenHamming; index++ {
		if a[index] != b[index] {
			errorsCount++
		}
	}

	return errorsCount, nil
}
