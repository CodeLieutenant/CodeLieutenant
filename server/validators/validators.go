package validators

import (
	"regexp"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func AlphaNumericUnicodeSpaceTranslationRegister(logger *zerolog.Logger, v *validator.Validate, trans ut.Translator) error {
	alphaNumericUnicodeSpace := regexp.MustCompile("^[\\p{L}\\p{N}\\s]+$")

	validate := func(field validator.FieldLevel) bool {
		return alphaNumericUnicodeSpace.MatchString(field.Field().String())
	}

	register := func(trans ut.Translator) error {
		return trans.Add("alphanumericunicodespace", "{0} must contain only alphanumeric characters including space", true)
	}

	translator := func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("alphanumericunicodespace", fe.Field())
		return t
	}

	if err := v.RegisterValidation("alphanumericunicodespace", validate); err != nil {
		return errors.Wrap(err, "Cannot register `alphanumericunicodespace` validator")
	}

	err := v.RegisterTranslation(
		"alphanumericunicodespace",
		trans,
		register,
		translator,
	)

	if err != nil {
		return errors.Wrap(err, "Cannot register `alphanumericunicodespace` translation")
	}

	return nil
}
