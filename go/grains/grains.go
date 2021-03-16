package grains

import (
	"errors"
)

// Total should return the total number of grains for chessboard (64 squares)
func Total() uint64 {
	grains, _ := Square(64)
	return grains*2 - 1
}

// Square should return the number of grains for a square position or an error
func Square(position int) (uint64, error) {
	if position < 1 || position > 64 {
		return 0, errors.New("square position should be between 1 and 64")
	}
	// Terms  in  the  Doubling  sequence  can  be  computed  by  the  function: d(n) = 2^n (for all natural numbers n)
	return 1 << (position - 1), nil
}
