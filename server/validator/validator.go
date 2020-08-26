package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	"log"
)

var (
	Uni       *ut.UniversalTranslator
	Validator *validator.Validate
	trans     ut.Translator
)

func ErrorHandler(ctx *fiber.Ctx, err error) {

	if err == gorm.ErrRecordNotFound {
		ctx.Status(404).JSON(struct {
			Message string `json:"message"`
		}{
			Message: "Record not found in database",
		})
	}

	if err, ok := err.(validator.ValidationErrors); ok {
		errors := map[string]string{}

		for _, valErr := range err {
			errors[valErr.Field()] = valErr.Translate(trans)
		}

		ctx.Status(422).JSON(struct {
			Errors map[string]string `json:"errors"`
		}{
			Errors: errors,
		})
		return
	}

	code := fiber.StatusInternalServerError

	// Check if it's an fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		ctx.Status(code).JSON(struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		})
		return
	}

	// Return HTTP response
	ctx.Status(code).JSON(struct {
		Message string `json:"message"`
	}{
		Message: "An error has occurred",
	})
}

func SetupValidator() {
	var found bool
	english := en.New()
	Uni = ut.New(english, english)
	trans, found = Uni.GetTranslator("en")
	if !found {
		log.Fatalln("Error translations not found")
	}

	Validator = validator.New()
	if err := Validator.RegisterValidation("alphanumericunicodespace", ValidateAlphaNumericUnicodeWithSpace); err != nil {
		log.Fatalf("Cannot register `alphanumericunicodespace` validator: %v", err)
	}

	if err := en_translations.RegisterDefaultTranslations(Validator, trans); err != nil {
		log.Fatalf("Cannot initialize validator: %v", err)
	}
}
