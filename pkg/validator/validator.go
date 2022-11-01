package validator

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
	"unicode"
)

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

func ValidatePasswordOzzo(pass string) error {
	err := validation.Validate(pass, validation.Required, validation.Match(regexp.MustCompile(`[a-zA-Z]`)))

	return err
}
