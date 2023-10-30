package validator

import (
	"errors"
	"fmt"
)

type IntValidators []intValidator

type intValidator interface {
	Validate(value int) error
}

type IntMaxValidator struct {
	Max int
}

func (v IntMaxValidator) Validate(value int) error {
	if value > v.Max {
		return fmt.Errorf("value must not be greater than %v", v.Max)
	}
	return nil
}

type IntMinValidator struct {
	Min int
}

func (v IntMinValidator) Validate(value int) error {
	if value < v.Min {
		return fmt.Errorf("value must not be greater than %v", v.Min)
	}
	return nil
}

func ValidateMapInt(name string, value map[string]any, rules IntValidators) (int, error) {
	rawValue, ok := value[name]
	if !ok {
		return -1, fmt.Errorf("missing key \"%v\"", name)
	}
	return ValidateInt(rawValue, rules)
}

func ValidateInt(value any, rules IntValidators) (int, error) {
	intValue, intOk := value.(int)
	floatValue, floatOk := value.(float64)
	if !intOk && !floatOk {
		return -1, errors.New("value is not a number")
	}

	if floatOk {
		if floatValue == float64(int(floatValue)) {
			intValue = (int(floatValue))
		} else {
			return -1, errors.New("value is not an int")
		}
	}

	for _, rule := range rules {
		if err := rule.Validate(intValue); err != nil {
			return -1, err
		}
	}

	return intValue, nil
}
