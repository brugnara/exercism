package leap

// IsLeapYear should return is a given year, is leap
func IsLeapYear(year int) bool {
	isDivisibleBy100 := (year % 100) == 0
	isDivisibleBy400 := (year % 400) == 0
	isDivisibleBy4 := (year % 4) == 0

	if isDivisibleBy4 {
		if isDivisibleBy100 {
			if isDivisibleBy400 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
