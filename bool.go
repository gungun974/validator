package validator

import (
	"errors"
	"fmt"
	"strings"
)

type BoolValidators []boolValidator

type boolValidator interface {
	Validate(value bool) error
}

type BoolIsTrueValidator struct{}

func (v BoolIsTrueValidator) Validate(value bool) error {
	if !value {
		return fmt.Errorf("value must be true")
	}
	return nil
}

type BoolIsFalseValidator struct{}

func (v BoolIsFalseValidator) Validate(value bool) error {
	if value {
		return fmt.Errorf("value must be false")
	}
	return nil
}

func ValidateMapBool(name string, value map[string]any, rules BoolValidators) (bool, error) {
	rawValue, ok := value[name]
	if !ok {
		return false, fmt.Errorf("missing key \"%v\"", name)
	}
	return ValidateBool(rawValue, rules)
}

func ValidateMapBoolOrFalse(
	name string,
	value map[string]any,
	rules BoolValidators,
) (bool, error) {
	rawValue, ok := value[name]
	if !ok {
		return false, nil
	}

	val, err := ValidateBool(rawValue, rules)

	return val, err
}

func ValidateBool(value any, rules BoolValidators) (bool, error) {
	boolValue, boolOk := value.(bool)
	intValue, intOk := value.(int)
	floatValue, floatOk := value.(float64)
	stringValue, stringOk := value.(string)
	if !boolOk && !intOk && !floatOk && !stringOk {
		return false, errors.New("value is not a bool")
	}

	if floatOk {
		if floatValue == float64(int(floatValue)) {
			intValue = (int(floatValue))
		} else {
			return false, errors.New("value is not a bool")
		}
	}

	if intOk || floatOk {
		if intValue == 1 {
			boolValue = true
		} else if intValue == 0 {
			boolValue = false
		} else {
			return false, errors.New("value is not a bool")
		}
	}

	if stringOk {
		normalizedStr := strings.ToLower(stringValue)

		if normalizedStr == "on" {
			boolValue = true
		} else if normalizedStr == "true" {
			boolValue = true
		} else if normalizedStr == "false" {
			boolValue = false
		} else {
			return false, errors.New("value is not a bool")
		}
	}

	for _, rule := range rules {
		if err := rule.Validate(boolValue); err != nil {
			return false, err
		}
	}

	return boolValue, nil
}
