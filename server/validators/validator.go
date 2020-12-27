package validators

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

func Register(logger *zerolog.Logger, v *validator.Validate, trans ut.Translator) (err error) {
	err = AlphaNumericUnicodeSpaceTranslationRegister(logger, v, trans)

	return
}
