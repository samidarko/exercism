package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) == len(b) {

		var errorCount int

		for i := 0; i < len(a); i++ {

			if a[i] != b[i] {
				errorCount++
			}

		}

		return errorCount, nil

	}

	return 0, errors.New("sequences are not equal")
}
