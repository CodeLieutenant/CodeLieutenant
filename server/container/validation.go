package container

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/malusev998/malusev998/server/validators"
)

func (c *Container) GetValidator() *validator.Validate {
	if c.validator == nil && c.translator == nil {
		english := en.New()
		uni := ut.New(english, english)

		c.translator, _ = uni.GetTranslator(c.Config.Locale)
		c.validator = validator.New()

		if err := en_translations.RegisterDefaultTranslations(c.validator, c.translator); err != nil {
			c.Logger.Fatal().Err(err).Msg("Error while registering english translations")
		}

		if err := validators.Register(c.validator, c.translator); err != nil {
			c.Logger.Fatal().Err(err).Msg("Error while registering custom validators")
		}
	}

	return c.validator
}

func (c *Container) GetTranslator() ut.Translator {
	if c.translator == nil {
		c.GetValidator()
	}

	return c.translator
}
