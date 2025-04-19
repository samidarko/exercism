package knapsack

type Item struct {
	Weight, Value int
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) int {
	// Crate DP table
	table := make([][]int, len(items)+1)
	for i := range table {
		table[i] = make([]int, maximumWeight+1)
	}

	for i := 1; i <= len(items); i++ {
		for j := 0; j <= maximumWeight; j++ {
			if items[i-1].Weight <= j {
				table[i][j] = Max(table[i-1][j], table[i-1][j-items[i-1].Weight]+items[i-1].Value)
			} else {
				table[i][j] = table[i-1][j]
			}
		}
	}
	return table[len(items)][maximumWeight]
}
