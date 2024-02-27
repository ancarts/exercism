
// Package leap provide a funclio that check if a sertain year is a leap year

package leap

// IsLeapYear returns true if the argument "year" is a leap year 
func IsLeapYear(year int) bool {

	return year%4 == 0 && year%100 != 0 || year%400 == 0

}
