package change

import "fmt"

// Change correctly determine the fewest number of coins to be given to a customer
func Change(coins []int, target int) ([]int, error) {
	if target < 0 {
		return nil, fmt.Errorf("error")
	}
	if target == 0 {
		return []int{}, nil
	}
	table := make([][]int, target+1, target+1)
	table[0] = make([]int, 0)

	for i := 0; i <= target; i++ {
		if table[i] != nil {
			for _, coin := range coins {
				currentSum := i + coin
				if currentSum < len(table) {

					// copy was necessary to avoid some references issue
					combinations := make([]int, len(table[i]))
					copy(combinations, table[i])
					combinations = append(combinations, coin)

					if table[currentSum] == nil || len(combinations) < len(table[currentSum]) {
						table[currentSum] = combinations
					}
				}
			}
		}
	}

	if len(table[target]) == 0 {
		return table[target], fmt.Errorf("error")
	}

	return table[target], nil
}
