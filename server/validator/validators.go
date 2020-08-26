package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var (
	alphaNumericUnicodeSpace = regexp.MustCompile("^[\\p{L}\\p{N}\\s]+$")
)

func ValidateAlphaNumericUnicodeWithSpace(fl validator.FieldLevel) bool {
	return alphaNumericUnicodeSpace.MatchString(fl.Field().String())
}
