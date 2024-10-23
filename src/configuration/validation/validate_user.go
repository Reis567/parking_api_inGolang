package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
	transl ut.Translator
)


func init () {
	if val,ok :=binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en,en)
		transl,_ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val,transl)
	}
}