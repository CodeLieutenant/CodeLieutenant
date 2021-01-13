package validators

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func Register(v *validator.Validate, trans ut.Translator) (err error) {
	err = AlphaNumericUnicodeSpaceTranslationRegister(v, trans)

	return
}
