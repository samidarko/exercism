package thefarm

import (
	"fmt"
)

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	if cows < 0 {
		return 0, fmt.Errorf("silly nephew, there cannot be %d cows", cows)
	}
	fodder, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction && fodder >= 0 {
		return fodder * 2 / float64(cows), nil
	}
	if err == ErrScaleMalfunction && fodder < 0 {
		return 0, fmt.Errorf("negative fodder")
	}
	if err == nil && fodder < 0 {
		return 0, fmt.Errorf("negative fodder")
	}
	if err != nil {
		return 0, err
	}

	return fodder / float64(cows), nil
}
