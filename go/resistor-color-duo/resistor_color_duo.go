package resistorcolorduo

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
		return -1
	}
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	return Color(colors[0])*10 + Color(colors[1])
}
