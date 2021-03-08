package grains

import (
	"errors"
	"math"
)

// Total should return the total number of grains for chessboard (64 squares)
func Total() uint64 {
	var totalGrains uint64
	for position := 1; position < 65; position++ {
		grains, _ := Square(position)
		totalGrains += grains
	}
	return totalGrains
}

// Square should return the number of grains for a square position or an error
func Square(position int) (uint64, error) {
	if position < 1 || position > 64 {
		return 0, errors.New("square position should be between 1 and 64")
	}
	return uint64(math.Pow(2.0, float64(position-1))), nil
}
