package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type FormValidator struct {
	validator *validator.Validate
}

func (fv *FormValidator) Validate(i interface{}) error {
	return fv.validator.Struct(i)
}

func NewFormValidator() *FormValidator {
	validate := validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &FormValidator{validate}
}

func ValidatorErrors(err error) map[string]string {
	fields := map[string]string{}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				fields[err.Field()] = fmt.Sprintf("field %s harus di isi", err.Field())
			default:
				fields[err.Field()] = fmt.Sprintf("%s error with tag %s should be %s", err.Field(), err.Tag(), err.Param())
			}
		}
	}

	return fields
}

func CheckDataEmpty(data ...any) error {
	for _, value := range data {
		if value == "" {
			return errors.New("field tidak boleh kosong")
		}
		if value == 0 {
			return errors.New("field tidak boleh kosong")
		}
	}
	return nil
}

func EmailFormat(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(email) {
		return nil
	}
	return errors.New("email tidak valid")
}

func MinLength(data string, minLength int) error {
	if len(data) < minLength {
		return errors.New("minimal " + strconv.Itoa(minLength) + " karakter, ulangi kembali!")
	}
	return nil
}
