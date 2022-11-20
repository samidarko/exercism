package queenattack

import (
	"fmt"
)

type coordinate struct {
	x uint8
	y uint8
}

// CanQueenAttack returns true if Queen can attack
func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if whitePosition == blackPosition {
		return false, fmt.Errorf("two queens cannot have same position: %s", whitePosition)
	}
	white, err := positionToCoordinates(whitePosition)
	if err != nil {
		return false, err
	}

	black, err := positionToCoordinates(blackPosition)
	if err != nil {
		return false, err
	}

	if white.x == black.x || white.y == black.y {
		return true, nil
	}

	numerator := float32(black.y) - float32(white.y)
	denominator := float32(black.x) - float32(white.x)

	slope := numerator / denominator

	return slope == 1 || slope == -1, nil
}

func positionToCoordinates(position string) (coordinate, error) {
	if len(position) != 2 {
		return coordinate{0, 0}, fmt.Errorf("invalid position: %s", position)
	}

	if position[0] < 'a' || position[0] > 'h' {
		return coordinate{0, 0}, fmt.Errorf("invalid column: %c", position[0])
	}

	if position[1] < '1' || position[1] > '8' {
		return coordinate{0, 0}, fmt.Errorf("invalid row: %c", position[1])
	}

	return coordinate{
		x: position[0] - 'a',
		y: position[1] - '0',
	}, nil
}
