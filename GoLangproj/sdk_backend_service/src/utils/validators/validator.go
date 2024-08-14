package validator

import (
	"github.com/go-playground/validator/v10"
)

// Title of the struct tag.
const tagName = "validate"

type StructValidator struct {
	validate *validator.Validate
}

func NewStructValidator() *StructValidator {

	sv := &StructValidator{
		validate: validator.New(),
	}

	sv.validate.SetTagName(tagName)
	return sv
}

func (sv *StructValidator) Validate(s interface{}) error {
	err := sv.validate.Struct(s)
	if err != nil {
		return err
	}
	return nil
}
