package leap

// TestVersion is a test version.
const TestVersion = 1

// IsLeapYear returns true if a year is a leap year, otherwise false.
func IsLeapYear(year int) bool {
	if year < 0 || year%4 != 0 || year%100 == 0 && year%400 != 0 {
		return false
	}
	return true
}
