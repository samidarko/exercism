package resistorcolortrio

import (
	"fmt"
	"math"
)

type Unit string

const (
	Ohms     Unit = "ohms"
	KiloOhms Unit = "kiloohms"
	MegaOhms Unit = "megaohms"
	GigaOhms Unit = "gigaohms"
)

type Size = int

const (
	KILO Size = 1_000
	MEGA Size = 1_000_000
	GIGA Size = 1_000_000_000
)

func Color(color string) int {
	switch color {
	case "black":
		return 0
	case "brown":
		return 1
	case "red":
		return 2
	case "orange":
		return 3
	case "yellow":
		return 4
	case "green":
		return 5
	case "blue":
		return 6
	case "violet":
		return 7
	case "grey":
		return 8
	case "white":
		return 9
	default:
		panic(fmt.Sprintf("unknown color %s", color))
	}
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	return Color(colors[0])*10 + Color(colors[1])
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	value := Value(colors) * int(math.Pow(10, float64(Color(colors[2]))))

	switch {
	case value < KILO:
		return fmt.Sprintf("%d %s", value, Ohms)
	case value < MEGA:
		return fmt.Sprintf("%d %s", value/KILO, KiloOhms)
	case value < GIGA:
		return fmt.Sprintf("%d %s", value/MEGA, MegaOhms)
	default:
		return fmt.Sprintf("%d %s", value/GIGA, GigaOhms)
	}
}
