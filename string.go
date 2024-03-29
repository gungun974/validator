package validator

import (
	"errors"
	"fmt"
	"net/mail"
	"strconv"
	"unicode/utf8"

	"github.com/nyaruka/phonenumbers"
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
		return fmt.Errorf("value length must not be greater than %v", v.Max)
	}
	return nil
}

type StringMinValidator struct {
	Min int
}

func (v StringMinValidator) Validate(value string) error {
	if utf8.RuneCountInString(value) < v.Min {
		return fmt.Errorf("value length must not be greater than %v", v.Min)
	}
	return nil
}

type StringEmailValidator struct{}

func (v StringEmailValidator) Validate(value string) error {
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("value is not an email")
	}
	return nil
}

type StringPhoneValidator struct{}

func (v StringPhoneValidator) Validate(value string) error {
	if _, err := phonenumbers.Parse(value, ""); err != nil {
		return fmt.Errorf("value is not an international phone number")
	}
	return nil
}

func ValidateMapString(name string, value map[string]any, rules StringValidators) (string, error) {
	rawValue, ok := value[name]
	if !ok {
		return "", fmt.Errorf("missing key \"%v\"", name)
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
	intValue, intOk := value.(int)
	floatValue, floatOk := value.(float64)
	if !stringOk && !intOk && !floatOk {
		return "", errors.New("value is not a string")
	}

	if intOk {
		stringValue = strconv.Itoa(intValue)
	}

	if floatOk {
		stringValue = strconv.FormatFloat(floatValue, 'f', -1, 64)
	}

	for _, rule := range rules {
		if err := rule.Validate(stringValue); err != nil {
			return "", err
		}
	}

	return stringValue, nil
}
