package yacht

import "sort"

// Score returns the score
func Score(dice []int, category string) int {

	switch category {
	case "yacht":
		for first, i := dice[0], 1; i < len(dice); i++ {
			if dice[i] != first {
				return 0
			}
		}
		return 50
	case "ones", "twos", "threes", "fours", "fives", "sixes":
		return countNumbers(dice, category)
	case "four of a kind":
		return fourOfAKind(dice)
	case "full house":
		return fullHouse(dice)
	case "little straight":
		return straight(dice, 1)
	case "big straight":
		return straight(dice, 2)
	case "choice":
		sum := 0
		for _, v := range dice {
			sum += v
		}
		return sum
	default:
		return 0
	}
}

func countNumbers(dice []int, category string) (result int) {
	categoryValue := map[string]int{
		"ones":   1,
		"twos":   2,
		"threes": 3,
		"fours":  4,
		"fives":  5,
		"sixes":  6,
	}
	value := categoryValue[category]
	for _, v := range dice {
		if v == value {
			result += value
		}
	}
	return
}

func fullHouse(dice []int) (result int) {
	valueCount := getValueCount(dice)
	values := getValues(valueCount)
	if len(values) != 2 {
		return
	}
	for _, value := range values {
		if valueCount[value] < 2 {
			return
		}
		result += valueCount[value] * value
	}
	return
}

func fourOfAKind(dice []int) int {
	valueCount := getValueCount(dice)
	values := getValues(valueCount)
	for _, value := range values {
		if valueCount[value] >= 4 {
			return value * 4
		}
	}
	return 0
}

func straight(dice []int, start int) int {
	sort.Ints(dice)
	for i, v := range dice {
		if v != start+i {
			return 0
		}
	}
	return 30
}

func getValueCount(dice []int) map[int]int {
	valueCount := map[int]int{}
	for _, v := range dice {
		valueCount[v]++
	}
	return valueCount
}

func getValues(valueCount map[int]int) (values []int) {
	for value, _ := range valueCount {
		values = append(values, value)
	}
	return
}
