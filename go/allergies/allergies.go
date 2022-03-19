package allergies

var allergensCode = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

// Allergies returns a list of allergies given an allergy score
func Allergies(allergies uint) []string {
	result := make([]string, 0)
	for allergen := range allergensCode {
		if AllergicTo(allergies, allergen) {
			result = append(result, allergen)
		}
	}
	return result
}

// AllergicTo returns true if patient is allergic to allergen
func AllergicTo(allergies uint, allergen string) bool {
	code := allergensCode[allergen]
	return allergies&code == code
}
