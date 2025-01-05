package internal

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
)

func NewValidator() *CustomValidator {
	v := validator.New(validator.WithRequiredStructEnabled())
	v.RegisterCustomTypeFunc(ValidatePGXText, pgtype.Text{})

	return &CustomValidator{validator: v}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)

	return err
}

func ValidatePGXText(field reflect.Value) interface{} {
	if text, ok := field.Interface().(pgtype.Text); ok {
		if text.Valid {
			return text.String
		}
	}

	return nil
}
