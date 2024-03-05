package validators

import "github.com/go-playground/validator/v10"

type IValidator func(validator.FieldLevel) bool

func New(validators map[string]IValidator) *validator.Validate {
	v := validator.New()

	for key, value := range validators {
		var fnc func(validator.FieldLevel) bool = value
		err := v.RegisterValidation(key, fnc)

		if err != nil {
			panic(err)
		}
	}

	return v
}
