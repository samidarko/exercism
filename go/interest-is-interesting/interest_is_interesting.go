package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		// 3.213% for a negative balance (balance gets more negative).
		return 3.213
	case balance < 1000:
		// 0.5% for a positive balance less than 1000 dollars.
		return 0.5
	case balance < 5000:
		//1.621% for a positive balance greater or equal than 1000 dollars and less than 5000 dollars.
		return 1.621
	default:
		//2.475% for a positive balance greater or equal than 5000 dollars.
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	yearsCount := 0
	for {
		balance = AnnualBalanceUpdate(balance)
		yearsCount++
		if balance >= targetBalance {
			break
		}
	}

	return yearsCount
}
