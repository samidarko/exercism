package darts

const (
	innerCircle  = 1
	middleCircle = 25  // squared value (real distance 5 units)
	outerCircle  = 100 // squared value (real distance 10 units)
)

func Score(x, y float64) int {

	radius := x*x + y*y
	switch {
	case radius <= innerCircle:
		return 10
	case radius <= middleCircle:
		return 5
	case radius <= outerCircle:
		return 1
	default:
		return 0

	}
}
