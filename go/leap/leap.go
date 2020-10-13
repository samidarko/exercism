package leap

// IsLeapYear Given a year, report if it is a leap year.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			return false
		}
		return true
	}
	return false
}
