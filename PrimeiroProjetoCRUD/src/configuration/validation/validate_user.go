package validation

import (
	"encoding/json"
	"errors"

	"github.com/felipebs10/Estudando-GO/tree/main/PrimeiroProjetoCRUD/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var ( /*Se entendi corretamente, é nesse momento que ele instancia as variáveis*/
	Validate = validator.New()
	transl   ut.Translator
)

/*Essa função init é utilizada para iniciar o programa e verificar se está tudo ok para seguir.*/
func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)

	}
}

func ValidateUserError(
	validation_err error,
) *rest_err.RestErr {

	//Instanciando as variáveis nessa função
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	//If para validar se foi enviando o tipo da variável correto
	if errors.As(validation_err, &jsonErr) { //Não pode usar um objeto que não seja ponteiro, por isso o &jsonErr
		return rest_err.NewBadRequestError("Invalid field type")
		//UnmarshallError é quando ele tenta dar um json para objeto, mas dá erro na tipagem de algun campo.
	} else if errors.As(validation_err, &jsonValidationError) { //Não pode usar um objeto que não seja ponteiro, por isso o &jsonValidationError
		errorsCauses := []rest_err.Causes{}
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}

}
