package validator

import (
	"errors"
	"fmt"
	"net/mail"
	"unicode/utf8"

	"github.com/ttacon/libphonenumber"
)

type StringValidators []stringValidator

type stringValidator interface {
	Validate(value string) error
}

type StringMaxValidator struct {
	Max int
}

func (v StringMaxValidator) Validate(value string) error {
	if utf8.RuneCountInString(value) > v.Max {
		return errors.New(fmt.Sprintf("value length must not be greater than %v", v.Max))
	}
	return nil
}

type StringMinValidator struct {
	Min int
}

func (v StringMinValidator) Validate(value string) error {
	if utf8.RuneCountInString(value) < v.Min {
		return errors.New(fmt.Sprintf("value length must not be greater than %v", v.Min))
	}
	return nil
}

type StringEmailValidator struct{}

func (v StringEmailValidator) Validate(value string) error {
	if _, err := mail.ParseAddress(value); err != nil {
		return errors.New(fmt.Sprintf("value is not an email"))
	}
	return nil
}

type StringPhoneValidator struct{}

func (v StringPhoneValidator) Validate(value string) error {
	if _, err := libphonenumber.Parse(value, ""); err != nil {
		return errors.New(fmt.Sprintf("value is not an international phone number"))
	}
	return nil
}

func ValidateMapString(name string, value map[string]any, rules StringValidators) (string, error) {
	rawValue, ok := value[name]
	if !ok {
		return "", errors.New(fmt.Sprintf("missing key \"%v\"", name))
	}
	return ValidateString(rawValue, rules)
}

func ValidateMapStringOrNil(
	name string,
	value map[string]any,
	rules StringValidators,
) (*string, error) {
	rawValue, ok := value[name]
	if !ok {
		return nil, nil
	}

	val, err := ValidateString(rawValue, rules)

	return &val, err
}

func ValidateString(value any, rules StringValidators) (string, error) {
	stringValue, stringOk := value.(string)
	if !stringOk {
		return "", errors.New("value is not a string")
	}

	for _, rule := range rules {
		if err := rule.Validate(stringValue); err != nil {
			return "", err
		}
	}

	return stringValue, nil
}
