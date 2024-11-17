package validation

import (
	"encoding/json"
	"errors"
	"meu-novo-projeto/src/configuration/rest_err"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"regexp"
)

var (
	validate *validator.Validate
	transl ut.Translator
)


func validatePassword(fl validator.FieldLevel) bool {
    password := fl.Field().String()

    // Verificar se a senha tem pelo menos 8 caracteres
    if len(password) < 8 {
        return false
    }

    // Verificar se contém pelo menos uma letra maiúscula
    if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
        return false
    }

    // Verificar se contém pelo menos um número
    if !regexp.MustCompile(`\d`).MatchString(password) {
        return false
    }

    // Verificar se contém pelo menos um caractere especial
    if !regexp.MustCompile(`[@$!%*?&]`).MatchString(password) {
        return false
    }

    return true
}


func init() {
    if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
        validate = val

        // Configurar tradução para inglês
        en := en.New()
        unt := ut.New(en, en)
        transl, _ = unt.GetTranslator("en")
        en_translation.RegisterDefaultTranslations(validate, transl)

        // Registrar validação personalizada para o campo Password
        validate.RegisterValidation("password", validatePassword)
    }
}


func ValidateUserError(
	validation_err error,
) *rest_err.RestErr{
	var jsonErr *json.UnmarshalTypeError
	var jsonVaidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr){
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonVaidationError) {
		errorCauses := []rest_err.Causes{}

		for _,e := range validation_err.(validator.ValidationErrors){
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Some fields are invalid !",errorCauses)
	}else{
		return rest_err.NewBadRequestError("Error trying to convert fields ")
	}
}