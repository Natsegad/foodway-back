package validator

import "unicode"

func ValidatePassword(pass string) (lower, upper bool, len, numberCount, symbolCount int) {

	for _, c := range pass {
		if unicode.IsNumber(c) {
			numberCount++
		} else if unicode.IsSymbol(c) {
			symbolCount++
		} else if unicode.IsLower(c) {
			lower = true
		} else if unicode.IsUpper(c) {
			upper = true
		}

		len++
	}

	return
}
