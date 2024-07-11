package utils

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate *validator.Validate
	Trans    ut.Translator
)

func InitializeValidatorUniversalTranslator() {
	Validate = validator.New()

	english := en.New()
	uni := ut.New(english, english)
	trans, found := uni.GetTranslator("en")
	if !found {
		panic("translator not found")
	}
	err := enTranslations.RegisterDefaultTranslations(Validate, trans)
	if err != nil {
		panic(err)
	}
	Trans = trans
}

func TranslateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

func Validation(a any) []string {

	err := Validate.Struct(a)
	errs := TranslateError(err, Trans)

	if len(errs) > 0 {
		var errorMessages []string
		for _, ve := range errs {
			errorMessages = append(errorMessages, ve.Error())
		}
		return errorMessages
	}
	return nil
}
