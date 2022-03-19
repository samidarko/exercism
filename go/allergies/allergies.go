package allergies

func Allergies(allergies uint) []string {
	result := make([]string, 0)
	allergens := []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}
	for _, allergen := range allergens {
		if AllergicTo(allergies, allergen) {
			result = append(result, allergen)
		}
	}
	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	switch allergen {
	case "eggs":
		return allergies&1 == 1
	case "peanuts":
		return allergies&2 == 2
	case "shellfish":
		return allergies&4 == 4
	case "strawberries":
		return allergies&8 == 8
	case "tomatoes":
		return allergies&16 == 16
	case "chocolate":
		return allergies&32 == 32
	case "pollen":
		return allergies&64 == 64
	case "cats":
		return allergies&128 == 128
	default:
		return false
	}
}
