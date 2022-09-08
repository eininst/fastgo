package valid

import (
	"fastgo/internal/code"
	"fastgo/internal/common/serr"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func ValidateStruct(s interface{}) error {
	var errors []*ErrorResponse
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}

		return serr.NewServiceError(errors, code.ERROR_PARAMETER)
	}
	return nil
}
