package hamming

import "errors"

func Distance(a, b string) (int, error) {
	aLen := len(a)
	if aLen == len(b) {

		var errorCount int

		for i := 0; i < aLen; i++ {

			if a[i] != b[i] {
				errorCount++
			}

		}

		return errorCount, nil

	}

	return 0, errors.New("sequences are not equal")
}
