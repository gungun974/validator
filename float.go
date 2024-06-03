package validator

import (
	"errors"
	"fmt"
	"strconv"
)

type FloatValidators []floatValidator

type floatValidator interface {
	Validate(value float64) error
}

type FloatMaxValidator struct {
	Max float64
}

func (v FloatMaxValidator) Validate(value float64) error {
	if value > v.Max {
		return fmt.Errorf("value must not be greater than %v", v.Max)
	}
	return nil
}

type FloatMinValidator struct {
	Min float64
}

func (v FloatMinValidator) Validate(value float64) error {
	if value < v.Min {
		return fmt.Errorf("value must not be greater than %v", v.Min)
	}
	return nil
}

func ValidateMapFloat(name string, value map[string]any, rules FloatValidators) (float64, error) {
	rawValue, ok := value[name]
	if !ok {
		return -1, fmt.Errorf("missing key \"%v\"", name)
	}
	return ValidateFloat(rawValue, rules)
}

func ValidateFloat(value any, rules FloatValidators) (float64, error) {
	intValue, intOk := value.(int)
	floatValue, floatOk := value.(float64)
	if !intOk && !floatOk {
		return -1, errors.New("value is not a number")
	}

	if intOk {
		floatValue = float64(intValue)
	}

	for _, rule := range rules {
		if err := rule.Validate(floatValue); err != nil {
			return -1, err
		}
	}

	return floatValue, nil
}

func CoerceAndValidateMapFloat(
	name string,
	value map[string]any,
	rules FloatValidators,
) (float64, error) {
	rawValue, ok := value[name]
	if !ok {
		return -1, fmt.Errorf("missing key \"%v\"", name)
	}
	return CoerceAndValidateFloat(rawValue, rules)
}

func CoerceAndValidateFloat(value any, rules FloatValidators) (float64, error) {
	stringValue, stringOk := value.(string)
	if stringOk {
		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err == nil {
			return ValidateFloat(floatValue, rules)
		}
	}

	return ValidateFloat(value, rules)
}
