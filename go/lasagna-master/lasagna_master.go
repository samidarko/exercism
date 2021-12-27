package lasagna

// PreparationTime estimate the preparation time
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// Quantities compute the amounts of noodles and sauce needed
func Quantities(layers []string) (int, float64) {
	noodlesCount, sauceCount := 0, 0
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodlesCount++
		case "sauce":
			sauceCount++
		}
	}
	return noodlesCount * 50, float64(sauceCount) * 0.2
}

// AddSecretIngredient add the secret ingredient
func AddSecretIngredient(friendsList []string, myList []string) []string {
	return append(myList[:], friendsList[len(friendsList)-1])
}

// ScaleRecipe scale the recipe
func ScaleRecipe(quantities []float64, portions int) []float64 {

	p := float64(portions)
	scaledQuantities := make([]float64, len(quantities))

	for i := range quantities {
		scaledQuantities[i] = quantities[i] / 2 * p
	}

	return scaledQuantities
}
